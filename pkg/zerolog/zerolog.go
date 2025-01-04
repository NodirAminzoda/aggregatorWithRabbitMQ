package zerolog

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"

	"aggreagtor/internal/costants"
)

type Logger struct {
	log *zerolog.Logger
}

func InitLogger() *Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount).
		Logger()
	return &Logger{log: &logger}
}

func (l *Logger) InfoWithPrefix(r string, prefix string, arg ...interface{}) {
	l.log.Info().
		Str(constants.XRequestID, r).
		Str("description", prefix).
		Msg(fmt.Sprint(arg...))
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *Logger) WarnWithPrefix(r string, prefix string, arg ...interface{}) {
	l.log.Warn().
		Str(constants.XRequestID, r).
		Str("description", prefix).
		Msg(fmt.Sprint(arg...))
}

func (l *Logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *Logger) ErrorWithPrefix(r string, prefix string, arg ...interface{}) {
	l.log.Error().
		Str(constants.XRequestID, r).
		Str("description", prefix).
		Msg(fmt.Sprint(arg...))
}

func (l *Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l *Logger) DebugWithPrefix(r string, prefix string, arg ...interface{}) {
	l.log.Debug().
		Str(constants.XRequestID, r).
		Str("description", prefix).
		Msg(fmt.Sprint(arg...))
}

func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *Logger) Fatalf(arg ...interface{}) {
	l.log.Fatal().
		Msg(fmt.Sprint(arg...))
}

func (l *Logger) Fatal(msg string) {
	l.log.Fatal().Msg(msg)
}
