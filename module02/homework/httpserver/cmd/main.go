package main

import (
	"context"
	"github.com/kevincgh/cncamp/httpserver/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
1. 接收客户端 `request`，并将 `request` 中带的 `header` 写入 `response header`
2. 读取当前系统的环境变量中的` VERSION` 配置，并写入 `response header`
3. `Server` 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 `server` 端的标准输出
4. 当访问 `localhost/healthz` 时，应返回 200
*/
func main() {
	srv := server.NewHTTPServer(":80")

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("server program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("server program exit error:", err)
		return
	}
	log.Println("server program exit ok")
}
