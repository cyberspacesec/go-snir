package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var (
	// Logger is the global logger instance
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           log.InfoLevel,
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		ReportCaller:    false,
	})

	// SlogHandler is a slog.Handler that uses our logger
	SlogHandler = slog.New(&slogToCharm{})
)

// EnableDebug enables debug logging
func EnableDebug() {
	Logger.SetLevel(log.DebugLevel)
}

// EnableSilence silences all logging
func EnableSilence() {
	Logger.SetOutput(io.Discard)
}

// Debug logs a debug message
func Debug(msg string, args ...interface{}) {
	Logger.Debug(msg, args...)
}

// Info logs an info message
func Info(msg string, args ...interface{}) {
	Logger.Info(msg, args...)
}

// Warn logs a warning message
func Warn(msg string, args ...interface{}) {
	Logger.Warn(msg, args...)
}

// Error logs an error message
func Error(msg string, args ...interface{}) {
	Logger.Error(msg, args...)
}

// Fatal logs a fatal message and exits
func Fatal(msg string, args ...interface{}) {
	Logger.Fatal(msg, args...)
}

// slogToCharm is a slog.Handler that uses our logger
type slogToCharm struct{}

func (h *slogToCharm) Enabled(ctx context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelDebug:
		return Logger.GetLevel() <= log.DebugLevel
	case slog.LevelInfo:
		return Logger.GetLevel() <= log.InfoLevel
	case slog.LevelWarn:
		return Logger.GetLevel() <= log.WarnLevel
	case slog.LevelError:
		return Logger.GetLevel() <= log.ErrorLevel
	default:
		return true
	}
}

func (h *slogToCharm) Handle(ctx context.Context, r slog.Record) error {
	level := log.InfoLevel
	switch r.Level {
	case slog.LevelDebug:
		level = log.DebugLevel
	case slog.LevelInfo:
		level = log.InfoLevel
	case slog.LevelWarn:
		level = log.WarnLevel
	case slog.LevelError:
		level = log.ErrorLevel
	}

	attrs := make([]interface{}, 0, r.NumAttrs()*2)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a.Key, a.Value.Any())
		return true
	})

	Logger.Log(level, r.Message, attrs...)
	return nil
}

func (h *slogToCharm) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *slogToCharm) WithGroup(name string) slog.Handler {
	return h
}