package main

import (
	"io/ioutil"
	"net/http"
)

var httpClient *http.Client

func getHttpClient() *http.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return httpClient
}

func setHttpClient(client *http.Client) {
	httpClient = client
}

func GetGoogle() ([]byte, error) {
	client := getHttpClient()
	resp, err := client.Get("https://google.com")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
