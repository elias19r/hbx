package basicerror

import (
	"fmt"
	"reflect"
)

type BasicError struct {
	code      int
	name      string
	message   string
	prevError error // A previous error, in case we need to extend an error.
}

func New(code int, name string) *BasicError {
	return &BasicError{
		code: code,
		name: name,
	}
}

// Code returns the error's code. We should use Code in other to compare errors.
// The code value 0 is reserved for ErrGeneral.
func (e *BasicError) Code() int {
	if e == nil {
		return 0
	}
	return e.code
}

func (e *BasicError) Name() string {
	if e == nil || e.code == 0 {
		return fmt.Sprintf("ErrGeneral(0)")
	}
	return fmt.Sprintf("%s(%d)", e.name, e.code)
}

func (e *BasicError) Message() string {
	if e == nil {
		return ""
	}
	return e.message
}

func (e *BasicError) PrevError() error {
	if e == nil {
		return nil
	}
	return e.prevError
}

func (e *BasicError) WithPrevError(err error) *BasicError {
	if e == nil {
		return &BasicError{}
	}
	return &BasicError{
		code:      e.code,
		name:      e.name,
		message:   e.message,
		prevError: err,
	}
}

// Error implements error interface.
func (e *BasicError) Error() string {
	if e == nil {
		return ""
	}
	str := e.Name()
	if e.message != "" {
		str += ": " + e.message
	}
	if reflect.ValueOf(e.prevError).Kind() == reflect.Ptr {
		if !reflect.ValueOf(e.prevError).IsNil() {
			str += ": " + e.prevError.Error()
		}
	} else if e.prevError != nil {
		str += ": " + e.prevError.Error()
	}
	return str
}

// String implements Stringer interface.
func (e *BasicError) String() string {
	return e.Error()
}

// WithMessage returns a new error based on e and with a custom message.
// Its parameters resemble fmt.Sprintf().
func (e *BasicError) WithMessage(format string, a ...interface{}) *BasicError {
	if e == nil {
		return &BasicError{}
	}
	return &BasicError{
		code:      e.code,
		name:      e.name,
		message:   fmt.Sprintf(format, a...),
		prevError: e.prevError,
	}
}
