package middleware

import (
	"log"
	"net/http"
	"os"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

func HeaderProcess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for key, values := range req.Header {
			for _, v := range values {
				w.Header().Set(key, v)
			}
		}
		w.Header().Set("VERSION", os.Getenv("VERSION"))
		next.ServeHTTP(w, req)
	})
}
