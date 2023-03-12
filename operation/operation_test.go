package operation

import (
	assert "calculator/helpTest"
	"testing"
)

func Test(t *testing.T) {
	sum := Operations["+"](2, 4)
	div := Operations["/"](4, 2)
	mult := Operations["*"](2, 4)
	sub := Operations["-"](2, 4)

	assert.Equals(6, sum, t)
	assert.Equals(2, div, t)
	assert.Equals(8, mult, t)
	assert.Equals(-2, sub, t)
}
