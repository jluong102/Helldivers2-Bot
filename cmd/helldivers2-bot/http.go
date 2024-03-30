package main

import (
	"net/http"
)

func SetJson(request *http.Request) {
	request.Header["Content-Type"] = []string{"application/json"}
}
