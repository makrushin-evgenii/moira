package worker

import (
	"time"

	"github.com/moira-alert/moira"
	"github.com/patrickmn/go-cache"
)

func (worker *Checker) newMetricsHandler(metricEventsChannel <-chan *moira.MetricEvent) error {
	for {
		metricEvent, ok := <-metricEventsChannel
		if !ok {
			return nil
		}
		pattern := metricEvent.Pattern
		if worker.needHandlePattern(pattern) {
			if err := worker.handleMetricEvent(pattern); err != nil {
				worker.Logger.Error().
					Error(err).
					Msg("Failed to handle metricEvent")
			}
		}
	}
}

func (worker *Checker) needHandlePattern(pattern string) bool {
	err := worker.PatternCache.Add(pattern, true, cache.DefaultExpiration)
	return err == nil
}

func (worker *Checker) handleMetricEvent(pattern string) error {
	start := time.Now()
	defer worker.Metrics.MetricEventsHandleTime.UpdateSince(start)
	worker.lastData = time.Now().UTC().Unix()
	triggerIds, err := worker.Database.GetPatternTriggerIDs(pattern)
	if err != nil {
		return err
	}
	// Cleanup pattern and its metrics if this pattern doesn't match to any trigger
	if len(triggerIds) == 0 {
		if err := worker.Database.RemovePatternWithMetrics(pattern); err != nil {
			return err
		}
	}
	worker.addTriggerIDsIfNeeded(triggerIds)
	return nil
}

func (worker *Checker) addTriggerIDsIfNeeded(triggerIDs []string) {
	needToCheckTriggerIDs := worker.getTriggerIDsToCheck(triggerIDs)
	if len(needToCheckTriggerIDs) > 0 {
		worker.Database.AddLocalTriggersToCheck(needToCheckTriggerIDs) //nolint
	}
}

func (worker *Checker) addRemoteTriggerIDsIfNeeded(triggerIDs []string) {
	needToCheckRemoteTriggerIDs := worker.getTriggerIDsToCheck(triggerIDs)
	if len(needToCheckRemoteTriggerIDs) > 0 {
		worker.Database.AddRemoteTriggersToCheck(needToCheckRemoteTriggerIDs) //nolint
	}
}

func (worker *Checker) getTriggerIDsToCheck(triggerIDs []string) []string {
	lazyTriggerIDs := worker.lazyTriggerIDs.Load().(map[string]bool)
	var triggerIDsToCheck []string = make([]string, 0, len(triggerIDs))
	for _, triggerID := range triggerIDs {
		if _, ok := lazyTriggerIDs[triggerID]; ok {
			randomDuration := worker.getRandomLazyCacheDuration()
			if err := worker.LazyTriggersCache.Add(triggerID, true, randomDuration); err != nil {
				continue
			}
		}
		if err := worker.TriggerCache.Add(triggerID, true, cache.DefaultExpiration); err == nil {
			triggerIDsToCheck = append(triggerIDsToCheck, triggerID)
		}
	}
	return triggerIDsToCheck
}
