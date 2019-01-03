package order

import "fmt"

type State int

const (
	StateNone      State = 0
	StateSubmitted State = 1
	StateRejected  State = 2
	StateActive    State = 3
	StateCancelled State = 4
	StateExpired   State = 5
	StateDone      State = 6
)

var stateStrings = map[State]string{
	StateNone:      "StateNone",
	StateSubmitted: "StateSubmitted",
	StateRejected:  "StateRejected",
	StateActive:    "StateActive",
	StateCancelled: "StateCancelled",
	StateExpired:   "StateExpired",
	StateDone:      "StateDone",
}

func (s State) String() string {
	if str := stateStrings[s]; str != "" {
		return str
	}
	return fmt.Sprintf("StateUnknown(%d)", int(s))
}
