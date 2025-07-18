package log

import (
	"log/slog"
	"testing"
)

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