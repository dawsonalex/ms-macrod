package config

import "log/slog"

type LogLevel string

const (
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
	LogLevelError LogLevel = "error"
	LogLevelWarn  LogLevel = "warn"
)

var logLevelToSlogMap = map[LogLevel]slog.Level{
	LogLevelInfo:  slog.LevelInfo,
	LogLevelDebug: slog.LevelDebug,
	LogLevelError: slog.LevelError,
	LogLevelWarn:  slog.LevelWarn,
}

type Log struct {
	Level LogLevel `ini:"level"`

	// HttpCorrelationIDHeaderKey points to a value sent for a client, for tracking across service boundaries.
	// This is the HTTP Header key value
	HttpCorrelationIDHeaderKey string `ini:"http_request_id_header_key"`

	// HttpCorrelationIDKey is the log key value of the cross-boundary ID.
	HttpCorrelationIDKey string `ini:"http_request_id_header_key"`

	// HTTP values to log.
	IncludeRequestBody    bool
	IncludeResponseBody   bool
	IncludeRequestHeader  bool
	IncludeResponseHeader bool
	IncludeUserAgent      bool
}

func (l LogLevel) ToSlogLevel() slog.Level {
	return logLevelToSlogMap[l]
}
