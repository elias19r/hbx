package numeric

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var (
	// regexpPattern specify the string representation of a Numeric.
	// Numeric values must be represented with a point and at least one unit
	// digit and one decimal digit. A negative sign is optional.
	// These vars must be considered constants.
	regexpPattern = `^(-?)(\d+)\.(\d{1,` + fmt.Sprintf("%d", Scale) + `})$`
	pattern       = regexp.MustCompile(regexpPattern)
)

// Numeric represents a decimal number as a big.Int.
type Numeric struct {
	value *big.Int
}

// New creates and a new Numeric and assigns it the value represented by the
// string s that must match the format specified by regexpPattern.
func New(s string) (Numeric, error) {
	// Match string and extract groups.
	if !pattern.MatchString(s) {
		return Numeric{}, ErrInvalidFormat("New(): regexp pattern mismatched: " + s)
	}

	// Here it must have 3 matching groups.
	matches := pattern.FindStringSubmatch(s)
	sign := matches[1]
	units := matches[2]
	decimals := matches[3]

	// Padding zeros to the right of decimal digits part so that we ensure
	// a correct Scale when creating the big.Int.
	decimals += strings.Repeat("0", Scale-len(decimals))

	// Try to create big.Int and return a new Numeric.
	value, ok := new(big.Int).SetString(units+decimals, 10)
	if !ok {
		// Must not reach here.
		return Numeric{}, ErrInvalidFormat("New(): could not create big.Int: " + s)
	}

	if sign == "-" {
		value.Neg(value)
	}

	return Numeric{value: value}, nil
}

// Copy makes a copy of n and returns a new Numeric.
func Copy(n Numeric) Numeric {
	if n.value == nil {
		return Numeric{value: new(big.Int)}
	}

	return Numeric{value: new(big.Int).Set(n.value)}
}
