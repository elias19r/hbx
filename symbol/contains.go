package symbol

import (
	"strings"
)

// Constants for the Contains function.
const (
	AsNone  = 0
	AsBase  = 1
	AsQuote = 2
)

// Contains returns: AsNone if s does not contains sym,
//                   AsBase if s contains sym as base asset,
//                   AsNone if s contains sym as quote asset.
func (s Symbol) Contains(sym Symbol) int {
	symbols := strings.Split(s.value, Separator)

	switch len(symbols) {
	case 1:
		if symbols[0] == sym.value {
			return AsBase
		}
	case 2:
		if symbols[0] == sym.value {
			return AsBase
		}
		if symbols[1] == sym.value {
			return AsQuote
		}
	}

	return AsNone
}
