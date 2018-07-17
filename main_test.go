package main_test

import (
	. "github.com/smartystreets/goconvey/convey"
	. "lander"
	"testing"
)

func TestGetConfig(t *testing.T) {
	Convey("Given an empty variable of type Config", t, func() {
		var config Config
		Convey("When we call the function GetConfig()", func() {
			config = GetConfig()
			Convey("config.Traefik shouldn't be blank", func() {
				So(config.Traefik, ShouldNotBeBlank)
			})
			Convey("config.Exposed shouldn't be blank", func() {
				So(config.Exposed, ShouldNotBeBlank)
			})
			Convey("config.Listen shouldn't be blank", func() {
				So(config.Listen, ShouldNotBeBlank)
			})
			Convey("config.Hostname should be blank", func() {
				So(config.Hostname, ShouldBeBlank)
			})
		})
	})
}

func TestPayloadDataGet(t *testing.T) {
	Convey("Given an empty and initialized variable of type PayloadData", t, func() {
		var payload = PayloadData{"", make(map[string][]Container)}
		Convey("And we call the method .Get()", func() {
			payload.Get()
			Convey("payload.Title should be blank", func() {
				So(payload.Title, ShouldBeBlank)
			})
			Convey("payload.Groups shouldn't be empty", func() {
				So(payload.Groups, ShouldNotBeEmpty)
			})
		})

	})
}
