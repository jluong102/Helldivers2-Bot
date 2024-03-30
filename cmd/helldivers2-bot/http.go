package main

import (
	"net/http"
)

// Make a generic HTTP request
func MakeRequest(request *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(request)
}

func SetJson(request *http.Request) {
	request.Header["Content-Type"] = []string{"application/json"}
}
