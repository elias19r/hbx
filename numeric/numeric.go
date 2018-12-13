package numeric

import (
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strings"
)

const (
	// SCALE determines the Numeric's scale. It is fixed to 36.
	// The maximum precision is not specified.
	SCALE = 36
)

var (
	// regexpPattern specify the string representation of a Numeric.
	// Numeric values must be represented with a point and at least one unit
	// digit and one decimal digit. A negative sign is optional.
	regexpPattern = `^(-?)(\d+)\.(\d{1,` + fmt.Sprintf("%d", SCALE) + `})$`
	pattern       = regexp.MustCompile(regexpPattern)
)

// Allocate some useful bigint.
var bigBlank = new(big.Int) // Blank bigint variable used when the calculated value can be discarded.
var bigRem = new(big.Int)   // Bigint to store remainder in QuoRem().
var bigZero = big.NewInt(0) // A always zero bigint.
var big10Pows []*big.Int    // Powers of ten from 10^0 to 10^SCALE

func init() {
	// Populate big10Pows.
	big10Pows = append(big10Pows, big.NewInt(1))  // 10^0
	big10Pows = append(big10Pows, big.NewInt(10)) // 10^1
	for i := 2; i <= SCALE; i++ {                 // 10^2 to 10^SCALE
		big10Pows = append(
			big10Pows,
			new(big.Int).Mul(big10Pows[1], big10Pows[i-1]),
		)
	}
}

// Numeric represents a decimal number as a bigint.
type Numeric struct {
	value *big.Int
}

// New creates and a new Numeric and assigns it the value represented by the
// string s that must match the format specified by regexpPattern.
func New(s string) (Numeric, error) {
	// Match string and extract groups.
	if !pattern.MatchString(s) {
		return Numeric{}, ErrInvalidFormat.WithMessage("New(): regexp pattern mismatched: %s", s)
	}
	matches := pattern.FindStringSubmatch(s)
	sign := matches[1]
	units := matches[2]
	decimals := matches[3]

	// Padding zeros to the right of decimal digits part so that we ensure
	// a correct scale when creating the bigint.
	decimals += strings.Repeat("0", SCALE-len(decimals))

	// Try to create bigint and return a new Numeric.
	value, ok := new(big.Int).SetString(units+decimals, 10)
	if !ok {
		// Must not reach here.
		return Numeric{}, ErrInvalidFormat.WithMessage("New(): could not create bigint: %s", s)
	}
	if sign == "-" {
		value.Neg(value)
	}
	return Numeric{value: value}, nil
}

// NewZero allocates a new Numeric with zero value.
func NewZero() Numeric {
	return Numeric{value: big.NewInt(0)}
}

// Copy makes a copy of n and returns a new Numeric.
func Copy(n Numeric) Numeric {
	if n.value == nil {
		log.Println(ErrNil.WithMessage("Copy(): n is nil"))
		return NewZero()
	}
	return Numeric{
		value: new(big.Int).Set(n.value),
	}
}

// Add sets c = a + b where c is a new Numeric.
func Add(a, b Numeric) Numeric {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Add(): a or b is nil"))
		return NewZero()
	}
	return Numeric{
		value: new(big.Int).Add(a.value, b.value),
	}
}

// Sub sets c = a - b where c is a new Numeric.
func Sub(a, b Numeric) Numeric {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Sub(): a or b is nil"))
		return NewZero()
	}
	return Numeric{
		value: new(big.Int).Sub(a.value, b.value),
	}
}

// Mul sets c = a * b where c is a new Numeric.
func Mul(a, b Numeric) Numeric {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Mul(): a or b is nil"))
		return NewZero()
	}
	c := Numeric{value: new(big.Int)}
	c.value.Mul(a.value, b.value)
	c.value.QuoRem(c.value, big10Pows[SCALE], bigBlank)
	return c
}

// Add sets a = a + b such that a value is overwritten.
func (a Numeric) Add(b Numeric) {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.Add(): a or b is nil"))
		return
	}
	a.value.Add(a.value, b.value)
}

// Sub sets a = a - b such that a value is overwritten.
func (a Numeric) Sub(b Numeric) {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.Sub(): a or b is nil"))
		return
	}
	a.value.Sub(a.value, b.value)
}

// Mul sets a = a * b such that a value is overwritten.
func (a Numeric) Mul(b Numeric) {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.Mul(): a or b is nil"))
		return
	}
	a.value.Mul(a.value, b.value)
	a.value.QuoRem(a.value, big10Pows[SCALE], bigBlank)
}

// Truncate limits a to n decimal digits.
func (a Numeric) Truncate(n int) {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.Truncate(): a is nil"))
		return
	}
	if n < 0 || n >= SCALE {
		return
	}
	n = SCALE - n
	bigBlank.QuoRem(a.value, big10Pows[n], bigRem)
	a.value.Sub(a.value, bigRem)
}

// RoundUp rounds a up at decimal digit n.
func (a Numeric) RoundUp(n int) {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.RoundUp(): a is nil"))
		return
	}
	if n < 0 || n >= SCALE {
		return
	}
	n = SCALE - n
	bigBlank.QuoRem(a.value, big10Pows[n], bigRem)
	if bigRem.Cmp(bigZero) == 1 {
		a.value.Sub(a.value, bigRem)
		a.value.Add(a.value, big10Pows[n])
	}
}

// IsNegative returns whether a < 0 or not.
func (a Numeric) IsNegative() bool {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.IsNegative(): a is nil"))
		return false
	}
	return a.value.Sign() == -1
}

// IsPositive returns whether a > 0 or not.
func (a Numeric) IsPositive() bool {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.IsPositive(): a is nil"))
		return false
	}
	return a.value.Sign() == 1
}

// IsZero returns whether a == 0 or not.
func (a Numeric) IsZero() bool {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.IsZero(): a is nil"))
		return false
	}
	return a.value.Sign() == 0
}

// LT stands for "less than" and compares a < b.
func (a Numeric) LT(b Numeric) bool {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.LT(): a or b is nil"))
		return false
	}
	return a.value.Cmp(b.value) == -1
}

// LTE stands for "less than or equal to" and compares a <= b.
func (a Numeric) LTE(b Numeric) bool {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.LTE(): a or b is nil"))
		return false
	}
	cmp := a.value.Cmp(b.value)
	return cmp == -1 || cmp == 0
}

// EQ stands for "equal" and compares a == b.
func (a Numeric) EQ(b Numeric) bool {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.EQ(): a or b is nil"))
		return false
	}
	return a.value.Cmp(b.value) == 0
}

// GT stands for "greater than" and compares a > b.
func (a Numeric) GT(b Numeric) bool {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.GT(): a or b is nil"))
		return false
	}
	return a.value.Cmp(b.value) == 1
}

// GTE stands for "greater than or equal to" and compares a >= b.
func (a Numeric) GTE(b Numeric) bool {
	if a.value == nil || b.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.GTE(): a or b is nil"))
		return false
	}
	cmp := a.value.Cmp(b.value)
	return cmp == 1 || cmp == 0
}

// string returns a representation of Numeric with all decimal digits.
func (a Numeric) string() string {
	// Padding zeros to ensure at least SCALE + 1 chars.
	str := a.value.String()
	str = strings.Repeat("0", SCALE-len(str)+1) + str

	// Reposition negative sign at beginning if any.
	if a.value.Sign() == -1 {
		str = "-" + strings.Replace(str, "-", "", 1)
	}

	// Place decimal point.
	return str[:len(str)-SCALE] + "." + str[len(str)-SCALE:]
}

// String returns a string representation of Numeric according to the format
// defined by regexpPattern.
func (a Numeric) String() string {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.String(): a is nil"))
		return "0.0"
	}
	return strings.TrimRight(a.string(), "0")
}

// StringWithScale returns a string representation of Numeric according to the
// format defined by regexpPattern and with a chosen scale n.
func (a Numeric) StringWithScale(n int) string {
	if a.value == nil {
		log.Println(ErrNil.WithMessage("Numeric.StringWithScale(): a is nil"))
		return "0.0"
	}
	if n < 1 || n > SCALE {
		n = SCALE
	}
	str := a.string()
	return str[:len(str)-SCALE+n]
}
