package main

import (
	"bytes"
	"github.com/clbanning/mxj"
	digest "github.com/xinsnake/go-http-digest-auth-client"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

// SOAP contains data for SOAP request
type SOAP struct {
	Body     string
	User     string
	Password string
	Address string
}

// SendRequest sends SOAP request to xAddr
func (soap SOAP) SendRequest(xaddr string) (mxj.Map, error) {
	// Create SOAP request
	request := soap.createRequest()
	// Make sure URL valid and add authentication in xAddr
	urlXAddr, err := url.Parse(xaddr)
	if err != nil {
		return nil, err
	}

	if soap.User != "" {
		urlXAddr.User = url.UserPassword(soap.User, soap.Password)
	}
	//glog.Info(request)
	// Create HTTP request
	buffer := bytes.NewBuffer([]byte(request))
	req, err := http.NewRequest("GET", urlXAddr.String(), buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/soap+xml")

	// Send request
	var httpClient = digest.NewRequest("admin", "Admin123", "GET", "http://admin:Admin123@192.168.1.64/ISAPI/System/capabilities", request)
	resp, err := httpClient.Execute()
	if err != nil {
		glog.Info(err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	glog.Infof("Hik response: %s", string(responseBody))


	return nil, nil
}

func (soap SOAP) createRequest() string {
	// Create request envelope
	request := `<?xml version="1.0" encoding="UTF-8"?>`


	// Clean request
	request = regexp.MustCompile(`\>\s+\<`).ReplaceAllString(request, "><")
	request = regexp.MustCompile(`\s+`).ReplaceAllString(request, " ")

	return request
}