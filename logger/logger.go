package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger - интерфейс для логгера. Позволяет использовать разные реализации.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	WithFields(fields logrus.Fields) *logrus.Entry
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

// Log - глобальная переменная для логгера.
var Log Logger

func init() {
	Log = newLogrusLogger()
}

type logrusLogger struct {
	*logrus.Logger
}

func newLogrusLogger() *logrusLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	logLevel := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
		logger.Warnf("Неизвестный уровень логирования: %s. Используется INFO.\n", logLevel)
	}
	logger.SetLevel(level)

	return &logrusLogger{logger}
}

// Debugf Реализация интерфейса Logger для logrusLogger
func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatalf(format, args...)
}

func (l *logrusLogger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

// Debug Добавленные методы для прямого логирования
func (l *logrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}
