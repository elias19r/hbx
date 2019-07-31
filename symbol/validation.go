package symbol

import "strings"

func validateValue(s string) error {
	if !pattern.MatchString(s) {
		return ErrInvalidFormat("validateValue(): regexp pattern mismatched: " + s)
	}

	return nil
}

func validateBaseQuote(s string) error {
	symbols := strings.Split(s, Separator)

	if len(symbols) == 2 && symbols[0] == symbols[1] {
		return ErrInvalidPair("validateBaseQuote(): base and quote are equal: " + s)
	}

	return nil
}
