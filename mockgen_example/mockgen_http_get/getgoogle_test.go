package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGoogle(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello, world!"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	// Replace http.Get with a function that returns the mock server's URL
	httpGet = func(url string) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Hello, world!")),
		}
		return resp, nil
	}

	// Call the function that makes an HTTP GET request
	response, err := GetGoogle()

	// Check for errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check the response body
	expectedString := "Hello, world!"
	if string(response) != expectedString {
		t.Errorf("Expected response body to contain %q, but got %q", expectedString, string(response))
	}
}
