package numeric

import "math/big"

// Sub sets c = a - b where c is a new Numeric.
func Sub(a, b Numeric) Numeric {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	return Numeric{value: new(big.Int).Sub(a.value, b.value)}
}

// Sub sets a = a - b such that a's value is overwritten and returns a.
func (a *Numeric) Sub(b Numeric) *Numeric {
	if a == nil {
		*a = Numeric{value: new(big.Int)}
	}
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	a.value.Sub(a.value, b.value)

	return a
}
