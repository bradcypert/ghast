package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testcase struct {
	Name   string
	Test   func(w http.ResponseWriter, r *http.Request)
	Status int
	Body   string
}

func TestResponses(t *testing.T) {

	controller := GhastController{}

	cases := []testcase{
		testcase{
			"200",
			func(w http.ResponseWriter, r *http.Request) { controller.Success(w, "foo") },
			200,
			"",
		},
		testcase{
			"403",
			func(w http.ResponseWriter, r *http.Request) { controller.Forbidden(w, "foo") },
			403,
			"",
		},
		testcase{
			"404",
			func(w http.ResponseWriter, r *http.Request) { controller.NotFound(w, "foo") },
			404,
			"",
		},
		testcase{
			"400",
			func(w http.ResponseWriter, r *http.Request) { controller.BadRequest(w, "foo") },
			400,
			"",
		},
		testcase{
			"401",
			func(w http.ResponseWriter, r *http.Request) { controller.Unauthorized(w, "foo") },
			401,
			"",
		},
		testcase{
			"500",
			func(w http.ResponseWriter, r *http.Request) { controller.InternalServerError(w, "foo") },
			500,
			"",
		},
	}

	for _, test := range cases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp := httptest.NewRecorder()
		test.Test(resp, req)

		if code := resp.Result().StatusCode; code != test.Status {
			t.Errorf("%s - Failed to set correct status code: %d", test.Name, code)
		}
	}
}
