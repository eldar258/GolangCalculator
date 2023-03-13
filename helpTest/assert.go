package assert

import (
	"testing"
)

func Equals(expected, result int, t *testing.T) {
	if expected != result {
		t.Errorf("%d != %d", expected, result)
	}
}

func EqualsString(expected, result string, t *testing.T) {
	if expected != result {
		t.Errorf("%s != %s", expected, result)
	}
}
