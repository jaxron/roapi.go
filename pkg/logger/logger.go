package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface defines the logging functionality required by the client.
type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
}

// NoOpLogger is a logger that does nothing, used as a default when no logger is provided.
type NoOpLogger struct{}

func (l *NoOpLogger) Debug(_ string, _ ...zapcore.Field) {}
func (l *NoOpLogger) Info(_ string, _ ...zapcore.Field)  {}
func (l *NoOpLogger) Warn(_ string, _ ...zapcore.Field)  {}
func (l *NoOpLogger) Error(_ string, _ ...zapcore.Field) {}

// DevelopmentLogger wraps a zap.Logger configured for development use.
type DevelopmentLogger struct {
	logger *zap.Logger
}

// NewDevelopmentLogger creates a logger that writes to stdout and includes a stack trace.
func NewDevelopmentLogger() Logger {
	// Configure the development logger
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return &DevelopmentLogger{logger: logger}
}

func (l *DevelopmentLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *DevelopmentLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *DevelopmentLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *DevelopmentLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

// ProductionLogger wraps a zap.Logger configured for production use.
type ProductionLogger struct {
	logger *zap.Logger
}

// NewProductionLogger creates a standard logger that writes to stdout.
func NewProductionLogger() Logger {
	// Configure the production logger
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return &ProductionLogger{logger: logger}
}

func (l *ProductionLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *ProductionLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ProductionLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *ProductionLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}
