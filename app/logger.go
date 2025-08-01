package app

import (
	"log/slog"
	"os"
	"path/filepath"
)

type Logger struct {
	file *os.File
	l    *slog.Logger
}

func newLogger(path string) (*Logger, error) {
	fileDir := filepath.Dir(path)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		err = os.MkdirAll(fileDir, 0o755)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, err
	}

	return &Logger{
		file: file,
		l:    slog.New(slog.NewTextHandler(file, nil)),
	}, nil
}

func (l *Logger) Error(msg string, args ...any) {
	l.l.Error(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.l.Info(msg, args...)
}

func (l *Logger) Close() {
	l.file.Close()
}

func SetupLogger() (*Logger, error) {
	errLogger, err := newLogger(errorDebugLogPath)
	if err != nil {
		return nil, err
	}

	errLog = errLogger

	return errLogger, nil
}

var errLog *Logger
