package address

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Address errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrNil = Error{basicerror.New(1, "ErrNil")}
)

func (e Error) Name() string {
	return "Address" + e.BasicError.Name()
}

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}

func (e Error) WithPrevError(err error) Error {
	return Error{e.BasicError.WithPrevError(err)}
}
