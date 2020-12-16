package middleware

import (
	"github.com/mattslocum/goserver/internal/logger"
	"net/http"
)


func HttpLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		logger.Info.Printf("%s %s", req.Method, req.URL.Path)
		next.ServeHTTP(rw, req)
	})
}
