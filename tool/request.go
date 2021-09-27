package tool

import (
	"github.com/go-resty/resty/v2"
	"log"
)

// Get request
func Get(url string) string {
	var err error
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		log.Panicf("the request is fail url %s and err message %s", url, err.Error())
	}
	return string(resp.Body())
}

// Post request
func Post(url string, data interface{}, contentType string) string {
	var err error
	client := resty.New()
	//client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.R().EnableTrace().SetBody(data).SetHeader("Content-Type", contentType).Post(url)
	if err != nil {
		log.Panicf("the request is fail url %s and err message %s", url, err.Error())
	}
	return string(resp.Body())
}
