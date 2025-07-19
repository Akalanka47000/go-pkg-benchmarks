package log

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)


func BenchmarkDiscardedLogZerolog(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	for b.Loop() {
		log.Debug().Msg("This is a discarded log")
	}
}

func BenchmarkSimpleInfoLogZerolog(b *testing.B) {
	for b.Loop() {
		log.Info().Msg("This is a simple info log")
	}
}

func BenchmarkInfoLogWithAttrsZerolog(b *testing.B) {
	for b.Loop() {
		log.Info().
			Str("key1", "value1").
			Int("key2", 42).
			Msg("This is an info log with attributes")
	}
}

func BenchmarkInfoLogWithGlobalAttrsZerolog(b *testing.B) {
	log.Logger = log.With().
		Str("application", "go-pkg-benchmarks").
		Int("version", 1).Logger()
	for b.Loop() {
		log.Info().Msg("This is an info log with global attributes")
	}
}

func BenchmarkInfoLogWithInheritedCtxZerolog(b *testing.B) {
	ctx := log.With().Str("trace", "123456").Logger().WithContext(b.Context())
	for b.Loop() {
		log.Ctx(ctx).Info().Msg("This is an info log with inherited context")
	}
}

func BenchmarkErrorLogZerolog(b *testing.B) {
	err := errors.New("This is the error itself")
	for b.Loop() {
		log.Error().Err(err).Msg("This is an error log")
	}
}

func BenchmarkErrorLogZerologWithStack(b *testing.B) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	err := errors.New("This is the error itself")
	for b.Loop() {
		log.Error().Stack().Err(err).Msg("This is an error log with stacktrace")
	}
}