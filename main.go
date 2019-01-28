package main

import (
	"github.com/golang/glog"

)

func main(){
	soap := SOAP{
		User:"admin",
		Password:"Admin123",
		Address: "http://admin:Admin123@192.168.1.64/ISAPI/System/capabilities",
	}
	re, _ := soap.SendRequest("http://admin:Admin123@192.168.1.64/ISAPI/System/capabilities")
	glog.Info(re)
}
