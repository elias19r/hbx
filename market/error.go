package market

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Market errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrNil            = Error{basicerror.New(1, "ErrNil")}
	ErrMarketNotFound = Error{basicerror.New(2, "ErrMarketNotFound")}
)

func (e Error) Name() string {
	return "Market" + e.BasicError.Name()
}

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}

func (e Error) WithPrevError(err error) Error {
	return Error{e.BasicError.WithPrevError(err)}
}
