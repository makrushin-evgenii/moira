package pagerduty

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/moira-alert/moira"
	mock_moira_alert "github.com/moira-alert/moira/mock/moira-alert"

	logging "github.com/moira-alert/moira/logging/zerolog_adapter"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSender(t *testing.T) {
	logger, _ := logging.ConfigureLog("stdout", "debug", "test", true)
	location, _ := time.LoadLocation("UTC")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	imageStore := mock_moira_alert.NewMockImageStore(mockCtrl)

	Convey("Init tests", t, func() {
		Convey("Has settings", func() {
			imageStore.EXPECT().IsEnabled().Return(true)
			senderSettings := map[string]string{
				"front_uri":   "http://moira.uri",
				"image_store": "s3",
			}
			imageStores := map[string]moira.ImageStore{
				"s3": imageStore,
			}
			sender := NewSender(senderSettings, logger, location, imageStores)
			So(sender.frontURI, ShouldResemble, "http://moira.uri")
			So(sender.logger, ShouldResemble, logger)
			So(sender.location, ShouldResemble, location)
			So(sender.imageStoreConfigured, ShouldResemble, true)
			So(sender.imageStore, ShouldResemble, imageStore)
		})
		Convey("Wrong image_store name", func() {
			senderSettings := map[string]string{
				"front_uri":   "http://moira.uri",
				"image_store": "s4",
			}
			imageStores := map[string]moira.ImageStore{
				"s3": imageStore,
			}
			sender := NewSender(senderSettings, logger, location, imageStores)
			So(sender, ShouldNotBeNil)
			So(sender.imageStoreConfigured, ShouldEqual, false)
			So(sender.imageStore, ShouldBeNil)
		})
		Convey("image store not configured", func() {
			imageStore.EXPECT().IsEnabled().Return(false)
			senderSettings := map[string]string{
				"front_uri":   "http://moira.uri",
				"image_store": "s3",
			}
			imageStores := map[string]moira.ImageStore{
				"s3": imageStore,
			}
			sender := NewSender(senderSettings, logger, location, imageStores)
			So(sender, ShouldNotBeNil)
			So(sender.imageStoreConfigured, ShouldEqual, false)
			So(sender.imageStore, ShouldBeNil)
		})
	})
}
