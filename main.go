package main

import (
	"context"
	"github.com/luoshengyue/coreDemo/framework/gin"
	"github.com/luoshengyue/coreDemo/framework/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	core := gin.New()

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr: ":8080",
	}
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<- quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}