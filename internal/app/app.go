package app

import (
	"aggreagtor/internal/handler"
	"aggreagtor/internal/modules/service"
	"aggreagtor/internal/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"aggreagtor/internal/config"
	"aggreagtor/internal/modules/adapter"
	"aggreagtor/pkg/zerolog"
)

type App struct {
	cfg    *config.Config
	log    *zerolog.Logger
	server *server.Server
}

func NewApp() *App {
	cfg := config.MustRead()
	log := zerolog.InitLogger()

	adap := adapter.New(&cfg.Adapter, log)
	servic := service.New(adap)
	hand := handler.New(&cfg.Server, log, servic)

	serv := server.New(cfg.Server, hand)

	return &App{
		cfg:    cfg,
		log:    log,
		server: serv,
	}
}

func (app *App) Run() {

	stopChanel := make(chan os.Signal, 1)
	signal.Notify(stopChanel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		app.log.Info("Input System start working")
		if err := app.server.Run(); err != nil {
			app.log.Fatalf("Could not listen on ", fmt.Sprintf("%s, %v", app.cfg.Server.Port, err))
		}
	}()

	<-stopChanel
	app.log.Info("Shutting down server gracefully ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.server.ShutDown(ctx); err != nil {
		app.log.Fatalf("server shutdown", fmt.Sprintf("%v", err))
	}

	app.log.Info("Server exited properly")
}
