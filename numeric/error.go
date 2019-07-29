package numeric

// ErrInvalidFormat occurs when an invalid string format is used to create a Numeric.
type ErrInvalidFormat string

func (e ErrInvalidFormat) Error() string {
	return "numeric.ErrInvalidFormat: " + string(e)
}
