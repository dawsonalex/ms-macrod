package requestid

import (
	"context"
	"github.com/google/uuid"
	sloghttp "github.com/samber/slog-http"
	"log/slog"
	"net/http"
)

// TODO: This needs tidying up to make testable

type ContextKey string

const (
	requestIdKey ContextKey = "request_id"
)

func AddToLogger(ctx context.Context, logger *slog.Logger) *slog.Logger {
	if id, ok := FromContext(ctx); ok {
		return logger.With(string(requestIdKey), id)
	}
	return logger
}

func AddToRequest(r *http.Request, id uuid.UUID) {
	sloghttp.AddCustomAttributes(r, slog.String(string(requestIdKey), id.String()))

	// TODO: we should probably add the id to the request context under out key here too.
}

func AddToContext(ctx context.Context, requestId uuid.UUID) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId.String())
}

func FromContext(ctx context.Context) (string, bool) {
	return getValue(ctx, requestIdKey)
}

func getValue(ctx context.Context, key ContextKey) (string, bool) {
	if val := ctx.Value(key); val != nil {
		return val.(string), true
	}

	return "", false
}
