package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/dawsonalex/ms-macrod/adapter/storage/inmemory"
	"github.com/dawsonalex/ms-macrod/config"
	"github.com/dawsonalex/ms-macrod/core/service"
	"github.com/dawsonalex/ms-macrod/httpserver"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func run(ctx context.Context, conf config.C) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := newLogger(conf.Log)
	logBuildInfo(logger)

	repo := inmemory.NewRepository()
	foodListingService, err := service.NewFoodListing(logger, repo)
	if err != nil {
		logger.Error(fmt.Sprintf("Fatal: cannot init food listing service: %v", err))
		os.Exit(1)
	}

	srv := httpserver.New(logger.WithGroup("http"), conf, foodListingService)
	httpServer := &http.Server{
		// TODO: Decide how to inject config here.
		Addr:    net.JoinHostPort(conf.Server.Host, conf.Server.Port),
		Handler: srv,
	}
	go func() {
		logger.Info("Starting HTTP Server", "address", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error(fmt.Sprintf("Error listening and serving: %s\n", err))
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Error(fmt.Sprintf("error shutting down http server: %s\n", err))
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()

	conf, err := config.ParseFile(config.FlagPath())
	if err != nil {
		l := newLogger(config.Default.Log)
		l.Warn("Error parsing config file, using defaults:")
		conf = config.Default
		_ = config.WriteTo(os.Stdout, *conf)
	}

	if err := run(ctx, *conf); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
