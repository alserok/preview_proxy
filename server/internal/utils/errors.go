package utils

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
)

func NewError(msg string, code uint) error {
	return &Err{
		msg:  msg,
		code: code,
	}
}
