package main

import (
	"context"
	"fmt"
	"github.com/slovty/mo-gin/internal/api"
	"github.com/slovty/mo-gin/pkg/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	injector, _, err := api.BuildInjector()
	if err != nil {
		return
	}
	r := route.InitRouter()
	srv := &http.Server{
		Addr:    "8080",
		Handler: r,
	}

	injector.Middleware.RegisterMiddleware(r)
	injector.Router.RegisterRouter(r)

	go func() {
		_ = r.Run()
	}()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		//todo send forced shutdown notification
		fmt.Println("Server forced to shutdown:", err)
	}

}
