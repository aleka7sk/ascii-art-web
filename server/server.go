package server

import (
	"ascii-art-web/config"
	apihttp "ascii-art-web/internal/api/delivery/http"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	HttpServer *http.Server
	Logger     *logrus.Logger
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(config config.Config) error {
	router := http.NewServeMux()
	apihttp.RegisterHTTPEndpoints(router)
	a.Logger = logrus.New()
	a.HttpServer = &http.Server{
		Addr:           ":" + config.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := a.HttpServer.ListenAndServe(); err != nil {
			a.Logger.Fatalf("Failed to listen and server: %+v", err)
		}
	}()
	a.Logger.Infof("Start Server on port: %v", config.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.HttpServer.Shutdown(ctx)
}
