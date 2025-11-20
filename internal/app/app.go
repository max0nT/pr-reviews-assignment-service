package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/max0nT/pr-assign/config"
	"github.com/max0nT/pr-assign/pkg/httpserver"
	"github.com/max0nT/pr-assign/pkg/logger"
	"github.com/max0nT/pr-assign/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New("debug")
	pg, err := postgres.New(cfg.PostgresUri)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Init http server
	httpServer := httpserver.New(l)

	httpServer.App.Run() // nolint: errcheck, gosec

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
