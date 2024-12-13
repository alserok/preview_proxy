package logger

import (
	"io"
	"log/slog"
	"os"
)

func newSlog(env string) *slogLogger {
	var (
		l      *slog.Logger
		output io.WriteCloser
	)

	switch env {
	case dev:
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prod:
		f, err := os.OpenFile("service.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("failed to open .logs file: " + err.Error())
		}

		l = slog.New(slog.NewTextHandler(io.MultiWriter(os.Stdout, f), &slog.HandlerOptions{Level: slog.LevelInfo}))
		output = f
	}

	return &slogLogger{
		l: l,
		f: output,
	}
}

type slogLogger struct {
	l *slog.Logger

	f io.WriteCloser
}

func (s *slogLogger) Close() error {
	return s.f.Close()
}

func (s *slogLogger) Info(msg string, v ...arg) {
	if len(v) == 0 {
		s.l.Info(msg)
	} else {
		args := make([]any, 0, len(v))
		for _, arg := range v {
			args = append(args, slog.Any(arg.key, arg.val))
		}

		s.l.Info(msg, args...)
	}
}

func (s *slogLogger) Debug(msg string, v ...arg) {
	if len(v) == 0 {
		s.l.Debug(msg)
	} else {
		args := make([]any, 0, len(v))
		for _, arg := range v {
			args = append(args, slog.Any(arg.key, arg.val))
		}

		s.l.Debug(msg, args...)
	}
}

func (s *slogLogger) Error(msg string, v ...arg) {
	if len(v) == 0 {
		s.l.Error(msg)
	} else {
		args := make([]any, 0, len(v))
		for _, arg := range v {
			args = append(args, slog.Any(arg.key, arg.val))
		}

		s.l.Error(msg, args...)
	}
}

func (s *slogLogger) Warn(msg string, v ...arg) {
	if len(v) == 0 {
		s.l.Warn(msg)
	} else {
		args := make([]any, 0, len(v))
		for _, arg := range v {
			args = append(args, slog.Any(arg.key, arg.val))
		}

		s.l.Warn(msg, args...)
	}
}
