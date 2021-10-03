package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestControllers(t *testing.T) {

	controller := GhastController{}

	var param []string

	cases := []struct {
		Name  string
		Test  func(w http.ResponseWriter, r *http.Request)
		Check func(actual []string) bool
		Body  string
	}{
		{
			"Query Param Exists",
			func(w http.ResponseWriter, r *http.Request) {
				param = controller.QueryParam(r, "name")
			},
			func(actual []string) bool {
				return "name" == actual[0]
			},
			"",
		},
		{
			"Query Param Does Not Exist",
			func(w http.ResponseWriter, r *http.Request) {
				param = controller.QueryParam(r, "no-name")
			},
			func(actual []string) bool {
				return len(actual) == 0
			},
			"",
		},
	}

	for _, test := range cases {
		req := httptest.NewRequest(http.MethodGet, "/?name=name", nil)
		resp := httptest.NewRecorder()
		test.Test(resp, req)

		if !test.Check(param) {
			t.Errorf("%s - Query Param unexpected value: %s", test.Name, param)
		}
	}
}
