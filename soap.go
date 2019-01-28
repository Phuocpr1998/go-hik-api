package main

import (
	"github.com/clbanning/mxj"
	"github.com/golang/glog"
	digest "github.com/xinsnake/go-http-digest-auth-client"
	"io/ioutil"
	"regexp"
)

// SOAP contains data for SOAP request
type SOAP struct {
	Body          string
	User          string
	Password      string
	Method        string
	Uri           string
	DeviceAddress string
}

// SendRequest sends SOAP request to xAddr
func (soap SOAP) SendRequest() (mxj.Map, error) {
	// Create SOAP request
	request := soap.createRequest()
	mapResponse := mxj.Map{}

	// Send request
	var httpClient = digest.NewRequest(soap.User, soap.Password, soap.Method, soap.DeviceAddress+soap.Uri, request)
	resp, err := httpClient.Execute()
	if err != nil {
		glog.Info(err)
		return mapResponse, err
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return mapResponse, err
	}

	glog.Infof("Hik response: %s", string(responseBody))

	mapResponse, err = mxj.NewMapXml(responseBody)
	if err != nil {
		return mapResponse, err
	}

	return mapResponse, nil
}

func (soap SOAP) createRequest() string {
	// Create request envelope
	request := `<?xml version="1.0" encoding="UTF-8"?>`

	// Clean request
	request = regexp.MustCompile(`\>\s+\<`).ReplaceAllString(request, "><")
	request = regexp.MustCompile(`\s+`).ReplaceAllString(request, " ")

	return request
}
