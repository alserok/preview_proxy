package utils

import (
	"context"
	"errors"
	"github.com/alserok/preview_proxy/server/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Err struct {
	msg  string
	code uint
}

func (e *Err) Error() string {
	return e.msg
}

const (
	Internal = iota
	BadRequest
	NotFound

	internal = "internal error"
)

func NewError(msg string, code uint) error {
	return &Err{
		msg:  msg,
		code: code,
	}
}

func ErrToGRPCError(ctx context.Context, in error) error {
	var (
		e *Err

		msg  string
		code codes.Code
	)

	if errors.As(in, &e) {
		switch e.code {
		case Internal:
			msg = internal
			code = codes.Internal
		case BadRequest:
			msg = e.msg
			code = codes.InvalidArgument
		case NotFound:
			msg = e.msg
			code = codes.NotFound
		default:
			msg = internal
			code = codes.Internal
		}
	} else {
		msg = internal
		code = codes.Internal
	}

	if code == codes.Internal {
		logger.FromContext(ctx).Error(msg)
	}

	return status.Error(code, msg)
}
