package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Client  *http.Client
	Request *http.Request
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
	}
}

func (client *HttpClient) Do(req *http.Request, successCallback func(response []byte),
	failCallback func(err error, code int, response []byte)) {
	if req == nil {
		failCallback(errors.New("request is miss"), -1, []byte{})
		return
	}
	response, err := client.Client.Do(req)
	defer response.Body.Close()
	if err != nil {
		failCallback(err, -1, []byte{})
		return
	}
	rawContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		failCallback(err, -1, []byte{})
		return
	}
	if response.StatusCode != http.StatusOK {
		failCallback(errors.New(fmt.Sprintf("unexpect statuscode, status code=%v",response.StatusCode)),
			response.StatusCode, rawContent)
		return
	}
	successCallback(rawContent)
}
