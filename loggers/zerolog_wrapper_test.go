package log

import (
	"testing"
)

func BenchmarkSimpleInfoLogZerologViaWrapper(b *testing.B) {
	for b.Loop() {
		Info("This is a simple info log called via wrapper")
	}
}

func BenchmarkInfoLogWithAttrsZerologViaWrapper(b *testing.B) {
	for b.Loop() {
		Info("This is an info log with attributes called via wrapper",
			"key1", "value1",
			"key2", 42,
		)
	}
}

func BenchmarkInfoLogWithInheritedCtxZerologViaWrapper(b *testing.B) {
	ctx := CtxWithFields(b.Context(), "trace", "123456")
	for b.Loop() {
		InfoCtx(ctx, "This is an info log with inherited context called via wrapper")
	}
}

func BenchmarkNopZeroLogViaWrapper(b *testing.B) {
	for b.Loop() {
		InfoCtx(CtxNop(b.Context()), "This is a nop log called via wrapper")
	}
}