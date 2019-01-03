package order

import "fmt"

type Side int

const (
	SideNone Side = 0
	SideBuy  Side = 1
	SideSell Side = 2
)

var sideStrings = map[Side]string{
	SideNone: "SideNone",
	SideBuy:  "SideBuy",
	SideSell: "SideSell",
}

func (s Side) String() string {
	if str := sideStrings[s]; str != "" {
		return str
	}
	return fmt.Sprintf("SideUnknown(%d)", int(s))
}
