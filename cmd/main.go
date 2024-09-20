package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dawsonalex/ms-macrod/build"
	"github.com/dawsonalex/ms-macrod/config"
	"github.com/dawsonalex/ms-macrod/httpserver"

	log "github.com/sirupsen/logrus"
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
	logger.SetFormatter(&log.TextFormatter{})

	b := build.Info()
	logger.WithFields(log.Fields{
		"version":     b.Version.Sprint(),
		"commit":      b.Commit,
		"branch":      b.Branch,
		"host":        b.Host,
		"environment": b.Environment,
	}).Info("Starting Server")

	confBytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		logger.Warn("Failed to marshal configuration")
	} else {
		logger.Infof("config:\n%s", confBytes)
	}

	srv := httpserver.New(conf)
	httpServer := &http.Server{
		// TODO: Decide how to inject config here.
		Addr:    net.JoinHostPort(conf.Server.Host, conf.Server.Port),
		Handler: srv,
	}
	go func() {
		logger.Infof("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
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
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func newLogger(conf config.Log) *log.Logger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(conf.Level)
	return logger
}

func main() {
	ctx := context.Background()

	conf, err := config.ParseFile(config.FlagPath())
	if err != nil {
		log.Warningf("error parsing config file, using defaults: %s", err)
		conf = config.Default
	}

	if err := run(ctx, *conf); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
