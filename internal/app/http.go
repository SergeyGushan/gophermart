package app

import (
	"fmt"
	"github.com/hanagantig/gracy"
	"go.uber.org/zap"
	"gophermart/internal/handler/http"
	"gophermart/internal/handler/http/api"
	client "net/http"
)

func (a *App) StartHTTPServer() error {
	go func() {
		a.startHTTPServer()
	}()

	err := gracy.Wait()
	if err != nil {
		a.logger.Error("failed to gracefully shutdown server", zap.Error(err))
		return err
	}
	a.logger.Info("server gracefully stopped")
	return nil
}

func (a *App) startHTTPServer() {
	handler := api.NewHandler(a.c.GetUseCase(), a.logger)

	router := http.NewRouter()
	router.
		WithHandler(handler, a.logger)

	srv := http.NewServer(a.cfg.RunAddress)
	srv.RegisterRoutes(router)

	gracy.AddCallback(func() error {
		return srv.Stop()
	})

	a.logger.Info(fmt.Sprintf("starting HTTP server at %s", a.cfg.RunAddress))
	err := srv.Start()
	if err != nil {
		a.logger.Fatal("Fail to start %s http server:", zap.Error(err))
	}
}

func (a *App) newHTTPClient() *client.Client {
	c := client.Client{}
	return &c
}
