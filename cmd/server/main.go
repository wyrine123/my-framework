package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"log/slog"
	v1 "my-framework/api/router/v1"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {
	// 加载配置
	// 初始化log配置
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)

	slog.SetDefault(logger)

	buildInfo, _ := debug.ReadBuildInfo()

	slog.Info("test", "build_info", buildInfo)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	rg := r.Group("/crm-core")
	{
		v1.ApiTestRouter(rg)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
		ErrorLog: slog.NewLogLogger(handler, slog.LevelInfo),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
