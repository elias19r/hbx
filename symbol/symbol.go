package symbol

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	// Separator determines the Symbol's separator character. It is set to "/".
	Separator = "/"

	// MinLen determines the minimum length for a symbol.
	MinLen = 3

	// MaxLen determines the maximum length for a symbol.
	MaxLen = 10
)

var (
	symbolFormat = `[A-Z]{` + fmt.Sprintf("%d", MinLen) + `,` + fmt.Sprintf("%d", MaxLen) + `}`
	// regexpPattern specify the string representation of a Symbol.
	regexpPattern = `^` + symbolFormat + `(` + Separator + symbolFormat + `)?$`
	pattern       = regexp.MustCompile(regexpPattern)
)

// Symbol identifies an asset or asset pair. e.g.: "BTC", "XRP", "XRP/BTC"
//
// NOTE: in case of an asset pair, for example "XRP/BTC", "XRP" is representing the base asset,
// and "BTC" is representing the quote asset.
type Symbol struct {
	value string
}

func New(s string) (Symbol, error) {
	if err := validateValue(s); err != nil {
		return Symbol{}, ErrInvalidFormat.WithMessage("New()").WithPrevError(err)
	}
	return Symbol{value: s}, nil
}

// EQ compares two symbols s == sym
func (s Symbol) EQ(sym Symbol) bool {
	return s.value == sym.value
}

// Contains returns: 0 if s does not contains sym,
//                   1 if s contains sym as base asset,
//                   2 if s contains sym as quote asset.
func (s Symbol) Contains(sym Symbol) int {
	symbols := strings.Split(s.value, Separator)
	// Here len(symbols) must equals 1 or 2.
	switch len(symbols) {
	case 1:
		if symbols[0] == sym.value {
			return 1
		}
	case 2:
		if symbols[0] == sym.value {
			return 1
		}
		if symbols[1] == sym.value {
			return 2
		}
	default:
		// Must not reach here.
		log.Println(ErrInvalidFormat.WithMessage("Contains(): invalid format for s: %s", s))
	}
	return 0
}

// Lowercase returns Symbol as string converted to lowercase.
func (s Symbol) Lowercase() string {
	return strings.ToLower(string(s.value))
}

// Strip returns Symbol as string with "/" removed, if any.
// e.g.: "XRP/BTC" returns "XRPBTC"
func (s Symbol) Strip() string {
	return strings.Replace(string(s.value), Separator, "", 1)
}

// Underscore returns Symbol as string with "/" replaced by "_", if any.
// e.g.: "XRP/BTC" returns "XRP_BTC"
func (s Symbol) Underscore() string {
	return strings.Replace(string(s.value), Separator, "_", 1)
}

func (s Symbol) Validate() error {
	return validateValue(s.value)
}

func validateValue(s string) error {
	// Match string pattern.
	if !pattern.MatchString(s) {
		return ErrInvalidFormat.WithMessage(
			"validateValue(): regexp pattern mismatched: %s", s,
		)
	}
	return nil
}
