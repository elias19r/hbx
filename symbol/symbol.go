package symbol

import (
	"fmt"
	"regexp"
)

const (
	// CharacterSet defines the allowed characters for a Symbol to have.
	CharacterSet = "[A-Z0-9]"

	// Separator determines the Symbol's separator character. It is set to "/".
	Separator = "/"

	// MinLen determines the minimum length for a Symbol.
	MinLen = 3

	// MaxLen determines the maximum length for a Symbol.
	MaxLen = 10
)

var (
	// These vars must be considered constants.
	symbolFormat  = CharacterSet + `{` + fmt.Sprintf("%d", MinLen) + `,` + fmt.Sprintf("%d", MaxLen) + `}`
	regexpPattern = `^` + symbolFormat + `(` + Separator + symbolFormat + `)?$`
	pattern       = regexp.MustCompile(regexpPattern)
)

// Symbol holds a string representation of an asset or asset pair symbol.
type Symbol struct {
	value string
}

// New returns a new valid Symbol based on string s.
func New(s string) (Symbol, error) {
	if err := validateValue(s); err != nil {
		return Symbol{}, err
	}
	if err := validateBaseQuote(s); err != nil {
		return Symbol{}, err
	}

	return Symbol{value: s}, nil
}
