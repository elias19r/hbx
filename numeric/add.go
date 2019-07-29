package numeric

import "math/big"

// Add sets c = a + b where c is a new Numeric.
func Add(a, b Numeric) Numeric {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	return Numeric{value: new(big.Int).Add(a.value, b.value)}
}

// Add sets a = a + b such that a's value is overwritten and returns a.
func (a *Numeric) Add(b Numeric) *Numeric {
	if a == nil {
		*a = Numeric{value: new(big.Int)}
	}
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	a.value.Add(a.value, b.value)

	return a
}
