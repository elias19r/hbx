package numeric

import "math/big"

// Truncate limits a to n decimal digits and returns a new Numeric.
// It returns an unmodified copy of a if n < 0 or n >= Scale.
func Truncate(a Numeric, n int) Numeric {
	if a.value == nil {
		a.value = new(big.Int)
	}

	// Make a copy.
	a = Numeric{value: new(big.Int).Set(a.value)}

	if n < 0 || n >= Scale {
		return a
	}

	n = Scale - n

	rem := new(big.Int)
	new(big.Int).QuoRem(a.value, pow10[n], rem)

	a.value.Sub(a.value, rem)

	return a
}

// Truncate limits a to n decimal digits such that a's value is overwritten
// and returns a.
// It returns a unmodified if n < 0 or n >= Scale.
func (a *Numeric) Truncate(n int) *Numeric {
	if n < 0 || n >= Scale {
		return a
	}

	if a == nil {
		*a = Numeric{value: new(big.Int)}
	}
	if a.value == nil {
		a.value = new(big.Int)
	}

	n = Scale - n

	rem := new(big.Int)
	new(big.Int).QuoRem(a.value, pow10[n], rem)

	a.value.Sub(a.value, rem)

	return a
}
