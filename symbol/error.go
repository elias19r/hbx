package symbol

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Symbol errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrInvalidFormat = Error{basicerror.New(1, "ErrInvalidFormat")}
)

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}
