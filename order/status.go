package order

import "fmt"

type Status int

const (
	StatusNone    Status = 0
	StatusNew     Status = 1
	StatusPartial Status = 2
	StatusFilled  Status = 3
)

var statusStrings = map[Status]string{
	StatusNone:    "StatusNone",
	StatusNew:     "StatusNew",
	StatusPartial: "StatusPartial",
	StatusFilled:  "StatusFilled",
}

func (s Status) String() string {
	if str := statusStrings[s]; str != "" {
		return str
	}
	return fmt.Sprintf("StatusUnknown(%d)", int(s))
}
