package tool

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Get request
func Get(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		log.Panicf("the request is fail url %s and err message %s", url, err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	result, _ := ioutil.ReadAll(response.Body)
	return string(result)
}

// Post request
func Post(url string, data interface{}, contentType string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonData, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Panicf("the request is fail url %s and err message %s", url, err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
