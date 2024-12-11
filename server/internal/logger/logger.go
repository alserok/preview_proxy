package logger

import "context"

type Logger interface {
	Info(msg string, v ...arg)
	Debug(msg string, v ...arg)
	Error(msg string, v ...arg)
	Warn(msg string, v ...arg)
}

const (
	Slog = iota
)

func NewLogger(t uint, env string) Logger {
	switch t {
	case Slog:
		return newSlog(env)
	default:
		panic("invalid logger type")
	}
}

func WithArg(key string, val any) arg {
	return arg{key: key, val: val}
}

type arg struct {
	key string
	val any
}

type ContextKey string

const (
	CtxLogger ContextKey = "logger"
)

func WrapLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, CtxLogger, logger)
}

func FromContext(ctx context.Context) Logger {
	logger, ok := ctx.Value(CtxLogger).(Logger)
	if !ok {
		panic("failed to extract logger from context")
	}

	return logger
}
