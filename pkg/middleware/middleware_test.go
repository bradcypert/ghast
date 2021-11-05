package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bradcypert/ghast/pkg/router"
)

func TestContextPassing(t *testing.T) {
	t.Run("should be able to get path variables", func(t *testing.T) {
		router := router.Router{}
		var contextVal string

		router.Get("/:name", func(w http.ResponseWriter, r *http.Request) {
			contextVal = r.Context().Value("FOO").(string)
		})

		router.AddMiddleware([]func(next http.Handler) http.Handler{
			func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					curContext := req.Context()
					req = req.Clone(context.WithValue(curContext, "FOO", "BAR"))

					next.ServeHTTP(w, req)
				})
			},
		})

		server := router.DefaultServer()
		req := httptest.NewRequest(http.MethodGet, "/foo", nil)
		resp := httptest.NewRecorder()
		server.Handler.ServeHTTP(resp, req)
		if contextVal != "BAR" {
			t.Error("Failed to set context as part of middleware chain")
		}
	})
}
