package logger

import "log/slog"

func newSlog(env string) *slogLogger {
	return &slogLogger{}
}

type slogLogger struct {
	l *slog.Logger
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
