package asset

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Asset errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrNil          = Error{basicerror.New(1, "ErrNil")}
	ErrInvalid      = Error{basicerror.New(2, "ErrInvalid")}
	ErrInvalidScale = Error{basicerror.New(3, "ErrInvalidScale")}
	ErrNotFound     = Error{basicerror.New(4, "ErrNotFound")}
)

func (e Error) Name() string {
	return "Asset" + e.BasicError.Name()
}

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}

func (e Error) WithPrevError(err error) Error {
	return Error{e.BasicError.WithPrevError(err)}
}
