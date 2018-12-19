package numeric

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Numeric errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrNil           = Error{basicerror.New(1, "ErrNil")}
	ErrInvalidFormat = Error{basicerror.New(2, "ErrInvalidFormat")}
)

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}
