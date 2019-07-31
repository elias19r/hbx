package symbol

// ErrInvalidFormat occurs when an invalid string format is used to create a Symbol.
type ErrInvalidFormat string

func (e ErrInvalidFormat) Error() string {
	return "symbol.ErrInvalidFormat: " + string(e)
}

// ErrInvalidPair occurs when a symbol is a pair of the same assets,
// i.e. base and quote are equal.
type ErrInvalidPair string

func (e ErrInvalidPair) Error() string {
	return "symbol.ErrInvalidPair: " + string(e)
}
