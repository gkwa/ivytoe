package ivytoe

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Main() int {
	fmt.Println("hello world")
	return 0
}

var noOpLogger = zap.NewNop()

type Logger struct {
	zap *zap.Logger
}

func Must(logger *Logger, err error) *Logger {
	if err != nil {
		panic(err)
	}
	return logger
}

func NewLogger(logFile string) *Logger {
	config := zap.NewProductionConfig()

	defaultLogLevel := zapcore.DebugLevel

	// config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeTime = nil

	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core)

	return &Logger{zap: logger}
}

func (l Logger) Debug(msg string, fields ...zap.Field) {
	l.writer().Debug(msg, fields...)
}

func (l Logger) Info(msg string, fields ...zap.Field) {
	l.writer().Info(msg, fields...)
}

func (l Logger) Error(msg string, fields ...zap.Field) {
	l.writer().Error(msg, fields...)
}

func (l Logger) Fatal(msg string, fields ...zap.Field) {
	l.writer().Error(msg, fields...)
}

func (l Logger) writer() *zap.Logger {
	if l.zap == nil {
		return noOpLogger
	}
	return l.zap
}
