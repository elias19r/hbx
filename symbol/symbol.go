package symbol

import (
	"strings"
)

// Symbol identifies an asset or asset pair. e.g.: "BTC", "XRP", "XRP/BTC"
type Symbol string

// Lowercase returns Symbol as string converted to lowercase.
func (s Symbol) Lowercase() string {
	return strings.ToLower(string(s))
}

// Strip returns Symbol as string with "/" removed, if any.
// e.g.: "XRP/BTC" returns "XRPBTC"
func (s Symbol) Strip() string {
	return strings.Replace(string(s), "/", "", 1)
}

// Underscore returns Symbol as string with "/" replaced by "_", if any.
// e.g.: "XRP/BTC" returns "XRP_BTC"
func (s Symbol) Underscore() string {
	return strings.Replace(string(s), "/", "_", 1)
}
