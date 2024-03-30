package main

import (
	"net/http"
)

func SetJson(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
}
