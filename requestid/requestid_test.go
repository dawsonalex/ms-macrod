package requestid

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"log/slog"
	"testing"
)

func TestAppendLoggableFields(t *testing.T) {
	ctx := context.Background()
	ctx = AddToContext(ctx, uuid.New())

	logOut := new(bytes.Buffer)
	logger := slog.New(slog.NewTextHandler(logOut, &slog.HandlerOptions{}))
	logger = AddToLogger(ctx, logger)

	logger.Info("test")
	t.Log(logOut.String())
}

func TestLogrusPackage(t *testing.T) {
	logger := log.StandardLogger()
	logger.Info("first comment")

	fieldsLogger := logger.WithFields(log.Fields{
		"test": "value",
	})

	fieldsLogger.Info("second comment")
	fieldsLogger.Logger.Info("third comment")
}
