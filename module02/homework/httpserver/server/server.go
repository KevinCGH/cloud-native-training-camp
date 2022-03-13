package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kevincgh/cncamp/httpserver/server/metrics"
	"github.com/kevincgh/cncamp/httpserver/server/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"net/http"
	"time"
)

type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(addr string) *HTTPServer {
	srv := &HTTPServer{
		srv: &http.Server{
			Addr: addr,
		},
	}
	metrics.Register()
	router := mux.NewRouter()
	router.HandleFunc("/", srv.rootHandler).Methods("GET")
	router.HandleFunc("/healthz", srv.healthzHandler).Methods("GET")
	router.HandleFunc("/latency", srv.latencyHandler).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	srv.srv.Handler = middleware.Logging(middleware.HeaderProcess(router))
	return srv
}

func (hs *HTTPServer) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		err = hs.srv.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	return hs.srv.Shutdown(ctx)
}

func (hs *HTTPServer) rootHandler(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello World")))
}

func (hs *HTTPServer) healthzHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (hs *HTTPServer) latencyHandler(w http.ResponseWriter, req *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randomInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randomInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d ms</h1>", randomInt)))
}
