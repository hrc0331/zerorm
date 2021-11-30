package logger

import (
	"context"
	"fmt"
	"time"
)

type Interface interface {
	SetLogMode(level LogLevel) Interface
	Info(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Trace(ctx context.Context, sql string, begin time.Time, err error)
	Printf(ctx context.Context, msg string)
}

func NewLogger(config Config) Interface {
	logger := logger{
		Config: Config{
			LogLevel: 		config.LogLevel,
			Color: 			config.Color,
			SlowThreshold: 	config.SlowThreshold,
		},
	}
	if len(logger.Color) == 0 {
		logger.Color = logLevel2Color[logger.LogLevel]
	}

	return &logger
}

type logger struct {
	Config
}

func (l *logger) SetLogMode(level LogLevel) Interface {
	newLogger := *l
	newLogger.LogLevel = level

	return &newLogger
}

func (l *logger) Info(ctx context.Context, msg string) {
	if l.LogLevel >= Info {
		l.Printf(ctx, msg)
	}
}

func (l *logger) Warn(ctx context.Context, msg string) {
	if l.LogLevel >= Warn {
		l.Printf(ctx, msg)
	}
}

func (l *logger) Error(ctx context.Context, msg string) {
	if l.LogLevel >= Error {
		l.Printf(ctx, msg)
	}
}

func (l *logger) Trace(ctx context.Context, sql string, begin time.Time, err error) {
	elapsedTime := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= Error:
		l.Printf(ctx, fmt.Sprintf("err msg: %s, sql: %s", err, sql))
	case elapsedTime > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
		l.Printf(ctx, fmt.Sprintf("slow sql with cost time >= %v, sql: %s", l.SlowThreshold, sql))
	case l.LogLevel >= Info:
		l.Printf(ctx, fmt.Sprintf("sql: %s", sql))
	default:
		return
	}
}

func (l *logger) Printf(ctx context.Context, msg string) {
	if _, ok := logLevel2Str[l.LogLevel]; ok {
		logLevel2Logger[l.LogLevel].Output(3, fmt.Sprintf(logLevel2Str[l.LogLevel], msg))
	}
}