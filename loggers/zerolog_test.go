package log

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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