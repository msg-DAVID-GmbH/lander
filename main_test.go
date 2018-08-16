package main_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "lander"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	dockerEnv := os.Getenv("LANDER_DOCKER")
	if dockerEnv != "" {
		Convey("Given an empty variable of type Config", t, func() {
			var RuntimeConfig Config
			Convey("When we call the function GetConfig()", func() {
				RuntimeConfig = GetConfig()
				Convey("RuntimeConfig.Traefik shouldn't be blank", func() {
					So(RuntimeConfig.Traefik, ShouldNotBeBlank)
				})
				Convey("RuntimeConfig.Exposed shouldn't be blank", func() {
					So(RuntimeConfig.Exposed, ShouldNotBeBlank)
				})
				Convey("RuntimeConfig.Listen shouldn't be blank", func() {
					So(RuntimeConfig.Listen, ShouldNotBeBlank)
				})
				Convey("RuntimeConfig.Hostname should be blank", func() {
					So(RuntimeConfig.Hostname, ShouldBeBlank)
				})
			})
		})
	}
}

//func TestPayloadDataGet(t *testing.T) {
//	Convey("Given an empty and initialized variable of type PayloadData", t, func() {
//		var payload = PayloadData{"", make(map[string][]Container)}
//		if _, err := os.Stat("/var/run/docker.sock"); os.IsNotExist(err) {
//			Convey("And we call the function GetContainer() while /var/run/docker.sock is not available, it should panic", func() {
//				So(GetContainers("unix:///var/run/docker.sock"), ShouldPanic)
//			})
//		} else {
//			Convey("And we call the method .Get() while /var/run/docker.sock is accessable", func() {
//				containers := GetContainers("unix:///var/run/docker.sock")
//				payload.Get(containers)
//				Convey("payload.Title should be blank", func() {
//					So(payload.Title, ShouldBeBlank)
//				})
//				Convey("payload.Groups shouldn't be empty", func() {
//					So(payload.Groups, ShouldNotBeEmpty)
//				})
//			})
//		}
//	})
//}
