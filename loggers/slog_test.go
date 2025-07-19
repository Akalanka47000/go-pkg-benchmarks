package log

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/pkg/errors"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func BenchmarkDiscardedLogSlog(b *testing.B) {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	for b.Loop() {
		slog.Debug("This is a discarded log")
	}
}

func BenchmarkSimpleInfoLogSlog(b *testing.B) {
	for b.Loop() {
		slog.Info("This is a simple info log")
	}
}

func BenchmarkInfoLogWithAttrsSlog(b *testing.B) {
	for b.Loop() {
		slog.Info("This is an info log with attributes",
			slog.String("key1", "value1"),
			slog.Int("key2", 42),
		)
	}
}

func BenchmarkInfoLogWithGlobalAttrsSlog(b *testing.B) {
	slog.SetDefault(slog.With(
		slog.String("application", "go-pkg-benchmarks"),
		slog.Int("version", 1),
	))
	for b.Loop() {
		slog.Info("This is an info log with global attributes")
	}
}

func BenchmarkInfoLogWithInheritedCtxSlog(b *testing.B) {
	key := "logger"
	ctx := context.WithValue(b.Context(), key, slog.With(
		slog.String("trace", "123456"),
	))
	logger := ctx.Value(key).(*slog.Logger)
	for b.Loop() {
		logger.InfoContext(b.Context(),"This is an info log with inherited context")
	}
}

func BenchmarkErrorLogSlog(b *testing.B) {
	err := errors.New("This is the error itself")
	for b.Loop() {
		slog.Error("This is an error log", slog.Any("error", err))
	}
}