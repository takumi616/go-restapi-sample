package web

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	Port    string
	Handler http.Handler
}

func (s *Server) Run(ctx context.Context) error {
	// Generate context with stopping signals
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Handler: s.Handler,
	}

	// Create http listener
	listener, err := net.Listen("tcp", ":"+s.Port)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return errors.New("failed to create http listener")
	}

	// Start server
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		err := server.Serve(listener)
		if err == http.ErrServerClosed {
			return nil
		} else {
			slog.ErrorContext(ctx, err.Error())
			return errors.New("failed to serve http")
		}
	})

	// Shutdown server
	<-ctx.Done()
	if err = server.Shutdown(context.Background()); err != nil {
		slog.ErrorContext(ctx, "failed to shut down http server")
	}

	// Wait for returned value from the goroutine
	if egErr := eg.Wait(); egErr != nil {
		return egErr
	}

	return err
}
