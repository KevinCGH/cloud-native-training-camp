package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kevincgh/cncamp/httpserver/server/middleware"
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

	router := mux.NewRouter()
	router.HandleFunc("/", srv.rootHandler).Methods("GET")
	router.HandleFunc("/healthz", srv.healthzHandler).Methods("GET")

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
