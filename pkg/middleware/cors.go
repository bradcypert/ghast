package middleware

import (
	"net/http"

	"github.com/bradcypert/ghast/pkg/router"
)

// CorsConfig struct for configuring CORS via the Cors middleware
type CorsConfig struct {
	AccessControlAllowOrigin  []string
	AccessControlAllowMethods []string
	AccessControlAllowHeaders []string
}

// Cors middleware
// Configures CORS headers on outgoing requests
func Cors(config CorsConfig) router.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			for _, item := range config.AccessControlAllowHeaders {
				rw.Header().Add("Access-Control-Allow-Headers", item)
			}
			for _, item := range config.AccessControlAllowMethods {
				rw.Header().Add("Access-Control-Allow-Methods", item)
			}
			for _, item := range config.AccessControlAllowOrigin {
				rw.Header().Add("Access-Control-Allow-Origin", item)
			}
			next.ServeHTTP(rw, req)
		})
	}

}
