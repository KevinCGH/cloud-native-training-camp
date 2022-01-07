package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
)

/**
1. 接收客户端 `request`，并将 `request` 中带的 `header` 写入 `response header`
2. 读取当前系统的环境变量中的` VERSION` 配置，并写入 `response header`
3. `Server` 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 `server` 端的标准输出
4. 当访问 `localhost/healthz` 时，应返回 200
*/
func main() {
	flag.Set("v", "4")
	glog.V(2).Infoln("Starting http server...")
	http.HandleFunc("/", logPanics(RootHandler))
	http.HandleFunc("/healthz", logPanics(HealthzHandler))

	// prof
	mux := http.NewServeMux()
	mux.HandleFunc("/", logPanics(RootHandler))
	mux.HandleFunc("/healthz", logPanics(HealthzHandler))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func RootHandler(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello World")))
}

func HealthzHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type HandleFnc func(w http.ResponseWriter, r *http.Request)

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				glog.Fatalf("[%v] caught panic: %v", request.RemoteAddr, x)

				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				glog.V(2).Infof("RemoteAddr: [%v], Method: %v, Status: %d", request.RemoteAddr, request.Method, http.StatusOK)
			}
		}()
		for key, values := range request.Header {
			for _, v := range values {
				writer.Header().Set(key, v)
			}
		}
		writer.Header().Set("VERSION", os.Getenv("VERSION"))
		function(writer, request)
	}
}
