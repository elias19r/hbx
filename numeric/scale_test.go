package numeric

import "testing"

func TestScale(t *testing.T) {
	t.Parallel()

	if Scale != 42 {
		t.Errorf("\nWant: %d\nHave: %d\n", 42, Scale)
	}
}
