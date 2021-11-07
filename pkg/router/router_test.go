package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPathParam(t *testing.T) {
	t.Run("should be able to get path variables", func(t *testing.T) {
		router := Router{}
		var name string

		router.Get("/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			name = router.PathParam(r, "name").(string)
		}))

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if name != "foo" {
			t.Error("Failed to set name via context params")
		}
	})
}

func TestNestingRouters(t *testing.T) {
	t.Run("should be able to nest routers", func(t *testing.T) {
		router := Router{}
		subrouter := Router{}

		var name string

		subrouter.Get("/:name", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				name = router.PathParam(r, "name").(string)
			},
		))

		router.Base("/v1").Merge(&subrouter)

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/v1/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if name != "foo" {
			t.Error("Failed to set name via context params")
		}
	})
}

func TestResponses(t *testing.T) {

	t.Run("should handle GETs correctly", func(t *testing.T) {
		router := Router{}

		router.Get("/:name", http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
				w.Write([]byte("hello"))
			},
		))

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Body.String() != "hello" {
			t.Error("Failed to set name via context params")
		}
	})

	t.Run("should set path variables", func(t *testing.T) {
		router := Router{}
		var name string

		router.Get("/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			name = r.Context().Value("name").(string)
		}))

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if name != "foo" {
			t.Error("Failed to set name via context params")
		}
	})

}
