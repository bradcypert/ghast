package middleware

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestLogger middleware
// Logs incoming requests by URL and body
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Incoming Request: " + req.URL.String())
		body, err := ioutil.ReadAll(req.Body)
		if err == nil && body != nil {
			fmt.Println(body)
		}

		if err != nil {
			fmt.Println("[Error]: Error when reading body on incoming request,", err)
		}

		next.ServeHTTP(rw, req)
	})
}
