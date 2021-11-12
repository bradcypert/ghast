package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradcypert/ghast/pkg/router"
)

func TestCorsMiddleware(t *testing.T) {
	t.Run("should set cors headers", func(t *testing.T) {
		router := router.Router{}

		router.Get("/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Yay"))
		}))

		router.AddMiddleware([]func(next http.Handler) http.Handler{
			Cors(CorsConfig{
				AccessControlAllowOrigin:  []string{"*"},
				AccessControlAllowMethods: []string{"GET"},
				AccessControlAllowHeaders: []string{"Content-Type"},
			}),
		})

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if resp.Result().Header["Access-Control-Allow-Origin"][0] != "*" {
			t.Error("Failed to set access control allow origin in cors middleware")
		}
		if resp.Result().Header["Access-Control-Allow-Methods"][0] != "GET" {
			t.Error("Failed to set access control allow origin in cors middleware")
		}
		if resp.Result().Header["Access-Control-Allow-Headers"][0] != "Content-Type" {
			t.Error("Failed to set access control allow origin in cors middleware")
		}
	})
}
