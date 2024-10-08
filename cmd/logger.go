package main

import (
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

func logBuildInfo(logger *slog.Logger) {
	b := build.Info()
	logger.With(
		"version", b.Version.Sprint(),
		"commit", b.Commit,
		"branch", b.Branch,
		"host", b.Host,
		"environment", b.Environment,
	).Info("Starting macrod")
}
