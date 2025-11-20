// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/max0nT/pr-assign/pkg/logger"
)

// Server -.
type Server struct {
	ctx context.Context
	eg  *errgroup.Group

	App    *gin.Engine
	notify chan error

	logger logger.Interface
}

// New -.
func New(l logger.Interface) *Server {
	group, ctx := errgroup.WithContext(context.Background())
	group.SetLimit(1) // Run only one goroutine

	s := &Server{
		ctx:    ctx,
		eg:     group,
		App:    nil,
		notify: make(chan error, 1),
		logger: l,
	}

	app := gin.Default()

	s.App = app

	return s
}

// Start -.
func (s *Server) Start() {
	s.eg.Go(func() error {
		err := s.App.Run()
		if err != nil {
			s.notify <- err

			close(s.notify)

			return err
		}

		return nil
	})

	s.logger.Info("http server - Server - Started")
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	var shutdownErrors []error

	// Wait for all goroutines to finish and get any error
	err := s.eg.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error(err, "http server - Server - Shutdown - s.eg.Wait")

		shutdownErrors = append(shutdownErrors, err)
	}

	s.logger.Info("http server - Server - Shutdown")

	return errors.Join(shutdownErrors...)
}
