package main

import (
	"encoding/json"
	"github.com/dawsonalex/ms-macrod/build"
	"github.com/dawsonalex/ms-macrod/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func newLogger(conf config.Log) *log.Logger {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetLevel(conf.Level)
	return logger
}

func logConfig(logger *log.Logger, conf config.C) {
	confBytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		logger.Warn("Failed to marshal configuration")
	} else {
		logger.Infof("config:\n%s", confBytes)
	}
}

func logBuildInfo(logger *log.Logger) {
	b := build.Info()
	logger.WithFields(log.Fields{
		"version":     b.Version.Sprint(),
		"commit":      b.Commit,
		"branch":      b.Branch,
		"host":        b.Host,
		"environment": b.Environment,
	}).Info("Starting Server")
}
