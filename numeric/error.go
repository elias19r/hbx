package numeric

import (
	"fmt"
)

type Error struct {
	code    int
	name    string
	message string
}

// List of Numeric errors. Code 0 is reserved for CustomError.
var (
	ErrNil           = &Error{1, "ErrNil", "value is nil"}
	ErrInvalidFormat = &Error{2, "ErrInvalidFormat", "invalid Numeric format"}
)

// Error implements error interface.
func (e *Error) Error() string {
	if e.message != "" {
		return e.Name() + ": " + e.message
	}
	return e.Name()
}

// String implements Stringer interface.
func (e *Error) String() string {
	return e.Error()
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Name() string {
	if e.name == "" {
		return fmt.Sprintf("CustomError(%d)", e.code)
	}
	return fmt.Sprintf("%s(%d)", e.name, e.code)
}

func (e *Error) Message() string {
	return e.message
}

// WithMessage returns a new error based on e and with a custom message.
// Its parameters resemble fmt.Sprintf().
func (e *Error) WithMessage(format string, a ...interface{}) *Error {
	return &Error{
		code:    e.code,
		name:    e.name,
		message: fmt.Sprintf(format, a...),
	}
}
