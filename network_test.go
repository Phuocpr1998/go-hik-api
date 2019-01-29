package hik_api

import (
	"testing"
)

func TestNetwork_EnableOnvif(t *testing.T) {
	device := Device{
		User:          "admin",
		Password:      "Admin123",
		DeviceAddress: "http://192.168.1.64",
	}

	err := device.EnableOnvif(true)
	if err != nil {
		t.Error(err)
	}
}
