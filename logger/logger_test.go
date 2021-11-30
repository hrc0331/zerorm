package logger

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestLogger_Info(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Info,
		SlowThreshold: time.Second,
	})
	logger.Info(context.Background(), "info")
}

func TestLogger_Warn(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Warn,
		SlowThreshold: time.Second,
	})
	logger.Warn(context.Background(), "warn")
}

func TestLogger_Error(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Error,
		SlowThreshold: time.Second,
	})
	logger.Error(context.Background(), "error")
}

func TestLogger_TraceInfo(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Info,
		SlowThreshold: time.Second,
	})
	logger.Trace(context.Background(), "should print", time.Now(), nil)
}

func TestLogger_TraceWarn(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Warn,
		SlowThreshold: time.Millisecond,
	})
	timeNow := time.Now()
	logger.Trace(context.Background(), "should not print", timeNow, nil)

	timeNow = time.Now()
	time.Sleep(time.Millisecond*2)
	logger.Trace(context.Background(), "should print", timeNow, nil)
}

func TestLogger_TraceError(t *testing.T) {
	logger := NewLogger(Config{
		LogLevel:      Error,
		SlowThreshold: time.Second,
	})
	logger.Trace(context.Background(), "should not print", time.Now(), nil)
	logger.Trace(context.Background(), "should print", time.Now(), errors.New("traceErr"))
}