package main

import (
	"io/ioutil"
	"net/http"
)

var httpGet = http.Get

func GetGoogle() ([]byte, error) {
	resp, err := httpGet("https://google.com")
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
