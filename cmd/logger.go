package main

import (
	"encoding/json"
	"fmt"
	"github.com/dawsonalex/ms-macrod/build"
	"github.com/dawsonalex/ms-macrod/config"
	"log/slog"
	"os"
)

func newLogger(conf config.Log) *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: conf.Level.ToSlogLevel(),
	})
	logger := slog.New(handler)
	return logger
}

func logConfig(logger *slog.Logger, conf config.C) {
	confBytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		logger.Warn("Failed to marshal configuration")
	} else {
		logger.Info(fmt.Sprintf("config:\n%s", confBytes))
	}
}

func logBuildInfo(logger *slog.Logger) {
	b := build.Info()
	logger.With(
		"version", b.Version.Sprint(),
		"commit", b.Commit,
		"branch", b.Branch,
		"host", b.Host,
		"environment", b.Environment,
	).Info("Starting Server")
}
