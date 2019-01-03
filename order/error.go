package order

import (
	"../basicerror"
)

type Error struct {
	*basicerror.BasicError
}

// List of Order errors. Each error must have a unique code.
// Code 0 is reserved for ErrGeneral.
var (
	ErrNil              = Error{basicerror.New(1, "ErrNil")}
	ErrNotFound         = Error{basicerror.New(2, "ErrNotFound")}
	ErrInvalidMarket    = Error{basicerror.New(3, "ErrInvalidMarket")}
	ErrInvalidMember    = Error{basicerror.New(4, "ErrInvalidMember")}
	ErrInvalidSide      = Error{basicerror.New(5, "ErrInvalidSide")}
	ErrInvalidType      = Error{basicerror.New(6, "ErrInvalidType")}
	ErrInvalidPrice     = Error{basicerror.New(7, "ErrInvalidPrice")}
	ErrInvalidAmount    = Error{basicerror.New(8, "ErrInvalidAmount")}
	ErrInvalidTimestamp = Error{basicerror.New(9, "ErrInvalidTimestamp")}
	ErrUnknownState     = Error{basicerror.New(10, "ErrUnknownState")}
)

func (e Error) Name() string {
	return "Order" + e.BasicError.Name()
}

func (e Error) WithMessage(format string, a ...interface{}) Error {
	return Error{e.BasicError.WithMessage(format, a...)}
}
