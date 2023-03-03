package main

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"io"
	"my/mockgen_http_get/mocks"
	"net/http"
	"testing"
)

func TestGetGoogle(t *testing.T) {
	// Replace http.Get with a function that returns the mocks server's URL

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockTrans := mocks.NewMockRoundTripper(mockCtl)
	req, _ := http.NewRequest(http.MethodGet, "https://google.com", nil)
	gomock.InOrder(
		mockTrans.EXPECT().RoundTrip(req).Return(
			&http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("Hello, world!")),
			},
			nil,
		),
	)

	setHttpClient(&http.Client{
		Transport: mockTrans,
	})

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
