package counter

import "testing"

func TestIncrement(t *testing.T) {
	Increment()

	if counter != 1 {
		t.Error("counter is expected to be incremented")
	}
}
