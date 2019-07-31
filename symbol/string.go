package symbol

import "strings"

func (s Symbol) String() string {
	return s.value
}

// Lowercase returns Symbol as string converted to lowercase.
func (s Symbol) Lowercase() string {
	return strings.ToLower(string(s.value))
}

// Strip returns Symbol as string with "/" removed, if any.
// e.g.: "BTC/USD" returns "BTCUSD"
func (s Symbol) Strip() string {
	return strings.Replace(string(s.value), Separator, "", 1)
}

// Underscore returns Symbol as string with "/" replaced by "_", if any.
// e.g.: "BTC/USD" returns "BTC_USD"
func (s Symbol) Underscore() string {
	return strings.Replace(string(s.value), Separator, "_", 1)
}
