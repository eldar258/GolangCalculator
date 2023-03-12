package assert

import "testing"

func Equals(expected int, result int, t *testing.T) {
	if expected != result {
		t.Errorf("%d != %d", expected, result)
	}
}
