package numeric

import (
	"math/big"
)

// Mul sets c = a * b where c is a new Numeric.
func Mul(a, b Numeric) Numeric {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	c := Numeric{value: new(big.Int)}

	c.value.Mul(a.value, b.value)
	c.value.QuoRem(c.value, pow10[Scale], new(big.Int))

	return c
}

// Mul sets a = a * b such that a's value is overwritten and returns a.
func (a *Numeric) Mul(b Numeric) *Numeric {
	if a == nil {
		*a = Numeric{value: new(big.Int)}
	}
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	a.value.Mul(a.value, b.value)
	a.value.QuoRem(a.value, pow10[Scale], new(big.Int))

	return a
}
