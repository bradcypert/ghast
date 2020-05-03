package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestLogger middleware
// Logs incoming requests by URL and body
func RequestLogger(rw *http.ResponseWriter, req *http.Request) {
	fmt.Println("Incoming Request: " + req.URL.String())
	body, err := ioutil.ReadAll(req.Body)
	if err == nil && body != nil {
		fmt.Println(body)
	}
}
