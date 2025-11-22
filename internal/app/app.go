package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/max0nT/pr-assign/internal/controllers"
	prrepo "github.com/max0nT/pr-assign/internal/repo/pull_request"
	teamrepo "github.com/max0nT/pr-assign/internal/repo/team"
	userrepo "github.com/max0nT/pr-assign/internal/repo/user"
	prmanage "github.com/max0nT/pr-assign/internal/usecase/pr_manage"
	teammanage "github.com/max0nT/pr-assign/internal/usecase/team_manage"

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

	userRepo := userrepo.New(pg)
	teamRepo := teamrepo.New(pg)
	prRepo := prrepo.New(pg)

	teamManage := teammanage.New(
		pg,
		userRepo,
		teamRepo,
	)
	prManage := prmanage.New(
		pg,
		userRepo,
		teamRepo,
		prRepo,
	)

	cnt := controllers.New(teamManage, prManage, validator.New())

	// Init http server
	httpServer := httpserver.New(l)
	group := httpServer.App.Group("/api/v1/")

	// Team manage
	group.POST("team/add/", cnt.AddTeam)
	group.GET("team/", cnt.GetTeam)

	// Pr manage
	group.GET("pr/", cnt.GetPr)
	group.POST("pr/open/", cnt.OpenPr)
	group.PATCH("pr/merge/", cnt.MergePr)
	group.PATCH("pr/reassign/", cnt.ReassignReviewer)

	// User manage
	group.PATCH("user/change-status-active/", cnt.ChangeUserActive)
	group.GET("user/", cnt.GetUsers)

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
