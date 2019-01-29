package hik_api

import (
	"github.com/golang/glog"
	"testing"
)

func TestEvents_EnableMotionEvent(t *testing.T) {
	device := Device{
		User:          "admin",
		Password:      "Admin123",
		DeviceAddress: "http://192.168.1.64",
	}

	err := device.EnableMotionEvent(true)
	if err != nil {
		glog.Info(err)
	}
}
