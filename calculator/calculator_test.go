package calculator

import (
	assert "calculator/helpTest"
	"testing"
)

func TestArabicDivision(t *testing.T) {
	expression := "2 / 5"
	result := "0"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)

	expression = "5 / 2"
	result = "2"

	expected = Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestArabicNegativeNum(t *testing.T) {
	expression := "2 - 5"
	result := "-3"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestRomanSum(t *testing.T) {
	expression := "III + IX"
	result := "XII"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestRomanDiv(t *testing.T) {
	expression := "X / II"
	result := "V"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestRomanMult(t *testing.T) {
	expression := "X * X"
	result := "C"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestRomanMult2(t *testing.T) {
	expression := "VI * VII"
	result := "XLII"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}

func TestRomanSup(t *testing.T) {
	expression := "X - I"
	result := "IX"

	expected := Calculate(expression)

	assert.EqualsString(expected, result, t)
}
