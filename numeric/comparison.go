package numeric

import "math/big"

// IsNegative returns whether a < 0 or not.
func (a Numeric) IsNegative() bool {
	if a.value == nil {
		return false
	}

	return a.value.Sign() == -1
}

// IsZero returns whether a == 0 or not.
func (a Numeric) IsZero() bool {
	if a.value == nil {
		return true
	}

	return a.value.Sign() == 0
}

// IsPositive returns whether a > 0 or not.
func (a Numeric) IsPositive() bool {
	if a.value == nil {
		return false
	}

	return a.value.Sign() == 1
}

// LT stands for "less than" and compares a < b.
func (a Numeric) LT(b Numeric) bool {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	return a.value.Cmp(b.value) == -1
}

// LTE stands for "less than or equal to" and compares a <= b.
func (a Numeric) LTE(b Numeric) bool {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	cmp := a.value.Cmp(b.value)

	return cmp == -1 || cmp == 0
}

// EQ stands for "equal" and compares a == b.
func (a Numeric) EQ(b Numeric) bool {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	return a.value.Cmp(b.value) == 0
}

// GT stands for "greater than" and compares a > b.
func (a Numeric) GT(b Numeric) bool {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	return a.value.Cmp(b.value) == 1
}

// GTE stands for "greater than or equal to" and compares a >= b.
func (a Numeric) GTE(b Numeric) bool {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if b.value == nil {
		b.value = new(big.Int)
	}

	cmp := a.value.Cmp(b.value)

	return cmp == 1 || cmp == 0
}
