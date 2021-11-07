package controllers

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
// Controllers only hvae the request exposed to them as
// the response writer is handled by the Ghast framework.
type RouteFunc func(req *http.Request) (Response, error)

// ServeHTTP calls f(w, r).
func (rf RouteFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// any error indicates an internal server error
	// need to make sure this is surfaced and clear in the docs
	response, err := rf(r)

	// Int's default type is 0, lets roll this up to a success
	if response.Status == 0 {
		response.Status = 200
	}

	if err != nil {
		w.WriteHeader(500)
	}

	// NEED TO FIGURE OUT HOW TO ADD HEADERS HERE
	if response.Body != nil {
		switch response.Body.(type) {
		case string:
			if err == nil {
				w.WriteHeader(response.Status)
			}
			w.Write([]byte(response.Body.(string)))
		case int:
			if err == nil {
				w.WriteHeader(response.Status)
			}
			w.Write([]byte(fmt.Sprint(response.Body.(int))))
		default:
			bytes, err := json.Marshal(response.Body)
			if err != nil {
				fmt.Println("ERR: Error when marshalling JSON passed from controller function")
				w.WriteHeader(500)
				w.Write(bytes)
			} else {
				if err == nil {
					w.WriteHeader(response.Status)
				}
				w.Write(bytes)
			}
		}
	}
}
