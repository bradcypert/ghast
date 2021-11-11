package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response a struct to define what a controller-based response looks like
type Response struct {
	Status  int
	Body    interface{}
	Headers http.Header
	Length  int
}

// RouteFunc a type alias for controller actions
// Controllers only have the request exposed to them as
// the response writer is handled by the Ghast framework.
type RouteFunc func(req *http.Request) (Response, error)

// ServeHTTP calls f(r).
func (rf RouteFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// any error indicates an internal server error
	// need to make sure this is surfaced and clear in the docs
	response, err := rf(r)

	var status int
	var body []byte

	for k, v := range response.Headers {
		for _, vals := range v {
			w.Header().Add(k, vals)
		}
	}

	if response.Body != nil {
		switch response.Body.(type) {
		case []byte:
			if err == nil {
				status = response.Status
			}
			body = response.Body.([]byte)
		case string:
			if err == nil {
				status = response.Status
			}
			body = []byte(response.Body.(string))
		case int:
			if err == nil {
				status = response.Status
			}
			body = []byte(fmt.Sprint(response.Body.(int)))
		default:
			bytes, mErr := json.Marshal(response.Body)
			if mErr != nil {
				fmt.Println("ERR: Error when marshalling JSON passed from controller function")
				status = 500
				body = bytes
			} else {
				if err == nil {
					w.Header().Set("Content-Type", "application/json")
					status = response.Status
				}
				body = bytes
			}
		}
	}

	if err != nil {
		status = 500
	}

	// Int's type default is 0, lets roll this up to a success
	if status == 0 {
		status = 200
	}

	w.WriteHeader(status)
	w.Write(body)
}
