package numeric

import (
	"math/big"
	"strings"
)

// String returns a string representation of Numeric according to the format
// defined by regexpPattern and with leading zeros trimmed.
func (a Numeric) String() string {
	if a.value == nil {
		a.value = new(big.Int)
	}

	return strings.TrimRight(a.string(), "0")
}

// StringWithScale returns a string representation of Numeric according to the
// format defined by regexpPattern and with a chosen scale n.
// If n is out of range, i.e. n < 1 or n > Scale, it sets n = Scale.
func (a Numeric) StringWithScale(n int) string {
	if a.value == nil {
		a.value = new(big.Int)
	}
	if n < 1 || n > Scale {
		n = Scale
	}

	str := a.string()

	return str[:len(str)-Scale+n]
}

// string returns a representation of Numeric with all decimal digits.
func (a Numeric) string() string {
	if a.value == nil {
		a.value = new(big.Int)
	}

	// Padding zeros to ensure at least Scale + 1 chars.
	str := a.value.String()
	str = strings.Repeat("0", Scale-len(str)+1) + str

	// Reposition negative sign at beginning if any.
	if a.value.Sign() == -1 {
		str = "-" + strings.Replace(str, "-", "", 1)
	}

	// Place decimal point.
	return str[:len(str)-Scale] + "." + str[len(str)-Scale:]
}
