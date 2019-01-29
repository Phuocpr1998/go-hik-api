package hik_api

import (
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
		t.Error(err)
	}
}
