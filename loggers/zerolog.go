package log

import (
	"context"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

// Ctx returns the Logger associated with the ctx. If no logger is associated, a disabled logger is returned.
var Ctx = zlog.Ctx

// With creates a child logger with the field added to its context.
var With = zlog.With

// Info logs an info message with the given fields.
func Info(message string, fields ...any) {
	zlog.Info().Fields(fields).Msg(message)
}

// InfoCtx logs an info message with the given fields and context.
func InfoCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Info().Fields(fields).Msg(message)
}

// Infof logs an info message with the given format and arguments.
func Infof(format string, args ...any) {
	zlog.Info().Msgf(format, args...)
}

// Error logs an error message with the given fields.
func Error(message string, fields ...any) {
	zlog.Error().Fields(fields).Msg(message)
}

// ErrorCtx logs an error message with the given fields and context.
func ErrorCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Error().Fields(fields).Msg(message)
}

// Errorf logs an error message with the given format and arguments.
func Errorf(format string, args ...any) {
	zlog.Error().Msgf(format, args...)
}

// Debug logs a debug message with the given fields.
func Debug(message string, fields ...any) {
	zlog.Debug().Fields(fields).Msg(message)
}

// DebugCtx logs a debug message with the given fields and context.
func DebugCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Debug().Fields(fields).Msg(message)
}

// Debugf logs a debug message with the given format and arguments.
func Debugf(format string, args ...any) {
	zlog.Debug().Msgf(format, args...)
}

// Warn logs a warning message with the given fields.
func Warn(message string, fields ...any) {
	zlog.Warn().Fields(fields).Msg(message)
}

// WarnCtx logs a warning message with the given fields and context.
func WarnCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Warn().Fields(fields).Msg(message)
}

// Warnf logs a warning message with the given format and arguments.
func Warnf(format string, args ...any) {
	zlog.Warn().Msgf(format, args...)
}

// Fatal logs a fatal message with the given fields and exits the program.
func Fatal(message string, fields ...any) {
	zlog.Fatal().Fields(fields).Msg(message)
}

// FatalCtx logs a fatal message with the given fields and context, then exits the program.
func FatalCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Fatal().Fields(fields).Msg(message)
}

// Fatalf logs a fatal message with the given format and arguments, then exits the program.
func Fatalf(format string, args ...any) {
	zlog.Fatal().Msgf(format, args...)
}

// Panic logs a panic message with the given fields and panics.
func Panic(message string, fields ...any) {
	zlog.Panic().Fields(fields).Msg(message)
}

// PanicCtx logs a panic message with the given fields and context, then panics.
func PanicCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Panic().Fields(fields).Msg(message)
}

// Panicf logs a panic message with the given format and arguments, then panics.
func Panicf(format string, args ...any) {
	zlog.Panic().Msgf(format, args...)
}

// Trace logs a trace message with the given fields.
func Trace(message string, fields ...any) {
	zlog.Trace().Fields(fields).Msg(message)
}

// TraceCtx logs a trace message with the given fields and context.
func TraceCtx(ctx context.Context, message string, fields ...any) {
	zlog.Ctx(ctx).Trace().Fields(fields).Msg(message)
}

// Tracef logs a trace message with the given format and arguments.
func Tracef(format string, args ...any) {
	zlog.Trace().Msgf(format, args...)
}

// WithField creates a child logger with the given field and its value added to its context.
func WithField(key string, value any) zerolog.Logger {
	return zlog.With().Any(key, value).Logger()
}

// WithFields creates a child logger with the field added to its context.
func WithFields(fields ...any) zerolog.Logger {
	return zlog.With().Fields(fields).Logger()
}

// CtxWithFields updates the context with a child logger that has the specified fields.
func CtxWithFields(ctx context.Context, fields ...any) context.Context {
	return WithFields(fields...).WithContext(ctx)
}

// Nop returns a context with a no-operation logger, effectively discarding all log messages.
func CtxNop(ctx context.Context) context.Context {
	return zerolog.Nop().WithContext(ctx)
}
