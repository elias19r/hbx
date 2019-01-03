package order

import "fmt"

type Type int

const (
	TypeNone   Type = 0
	TypeLimit  Type = 1
	TypeMarket Type = 2
	TypeStop   Type = 3
)

var typeStrings = map[Type]string{
	TypeNone:   "TypeNone",
	TypeLimit:  "TypeLimit",
	TypeMarket: "TypeMarket",
	TypeStop:   "TypeStop",
}

func (t Type) String() string {
	if str := typeStrings[t]; str != "" {
		return str
	}
	return fmt.Sprintf("TypeUnknown(%d)", int(t))
}
