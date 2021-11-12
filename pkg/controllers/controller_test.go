package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradcypert/ghast/pkg/router"
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
				return actual[0] == "name"
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

func TestControllerUnmarshalling(t *testing.T) {

	controller := GhastController{}
	jsonBytes, _ := json.Marshal(struct {
		Name string
		Age  int
	}{
		"Brad",
		28,
	})

	req := httptest.NewRequest(http.MethodPost, "/?name=name", bytes.NewReader(jsonBytes))

	var user struct {
		Name string
		Age  int
	}
	controller.Unmarshal(req, &user)
	if user.Name != "Brad" {
		t.Errorf("Got: \"%s\"; expected value: \"Brad\"", user.Name)
	}

	if user.Age != 28 {
		t.Errorf("Got: \"%d\"; expected value: 28", user.Age)
	}
}

type MockController struct {
	GhastController
}

func (m MockController) Index(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: "hello world",
	}, nil
}

func (m MockController) Get(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: "hello world",
	}, nil
}

func (m MockController) Create(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: "hello world",
	}, nil
}

func (m MockController) Update(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: "hello world",
	}, nil
}

func (m MockController) Delete(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: "hello world",
	}, nil
}

type MockResp struct {
	User string
}

func (m MockController) Struct(req *http.Request) (router.Response, error) {
	return router.Response{
		Body: MockResp{
			User: "Brad",
		},
	}, nil
}

func (m MockController) GetName() string {
	return "mock"
}

func TestRouterWorksWithControllers(t *testing.T) {

	t.Run("should handle controller response functions properly", func(t *testing.T) {
		r := router.Router{}

		controller := MockController{}

		r.Get("/:name", router.RouteFunc(controller.Get))

		server := r.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Body.String() != "hello world" {
			t.Error("Failed to set name via context params, got ", resp.Body)
		}
	})

	t.Run("passing a struct should marshall body and set content type", func(t *testing.T) {
		r := router.Router{}

		controller := MockController{}

		r.Get("/:name", router.RouteFunc(controller.Struct))

		server := r.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Body.String() != "{\"User\":\"Brad\"}" {
			t.Error("Failed to set name via marshalling, got ", resp.Body)
		}

		if resp.Result().Header["Content-Type"][0] != "application/json" {
			t.Error("Failed to set content type, got  ", resp.Result().Header["Content-Type"][0])
		}
	})
}

func TestResources(t *testing.T) {
	t.Run("Resources should work", func(t *testing.T) {
		r := router.Router{}

		controller := MockController{}

		r.Resource("/", controller)

		server := r.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/mock", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Result().StatusCode != 200 {
			t.Error("Failed to set status code properly ", resp.Result().StatusCode)
		}
		if resp.Body.String() != "hello world" {
			t.Error("Failed to set name via context params, got ", resp.Body)
		}
	})
}
