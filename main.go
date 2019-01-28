package main

import (
	"github.com/golang/glog"
)

func main() {
	soap := SOAP{
		User:          "admin",
		Password:      "Admin123",
		DeviceAddress: "http://192.168.1.64",
		Uri:           "/ISAPI/System/capabilities",
		Method:        "GET",
	}
	re, _ := soap.SendRequest()
	glog.Info(re)
}
