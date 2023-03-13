package calculator

import (
	"calculator/operation"
	"calculator/roman"
	"log"
	"strconv"
	"strings"
)

func Calculate(expression string) string {
	a, b, operator, isRoman := separate(expression)

	result := operation.Operations[operator](a, b)

	if isRoman {
		if result < 1 {
			log.Fatal("No Roman numerals less than 'I'!")
		}
		return roman.ToRoman(result)
	} else {
		return strconv.Itoa(result)
	}
}

func separate(expression string) (int, int, string, bool) {
	values := strings.Split(expression, " ")

	if values == nil || len(values) != 3 {
		log.Fatal("Wrong data format!", len(values))
	}

	aString := values[0]
	operator := values[1]
	bString := values[2]

	validateOperation(operator)
	var isRomanNum bool
	a, b := getNumbers(aString, bString, &isRomanNum)
	return a, b, operator, isRomanNum
}

func getNumbers(aString string, bString string, isRoman *bool) (int, int) {
	isARoman := roman.IsRomanNumOneToTen(aString)
	isBRoman := roman.IsRomanNumOneToTen(bString)

	if isARoman != isBRoman {
		log.Fatal("Numbers in different formats!")
	}
	*isRoman = isARoman

	if !isARoman {
		a := parseInt(aString)
		b := parseInt(bString)
		return a, b
	}

	return roman.NumOneToTen[aString], roman.NumOneToTen[bString]
}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("'%s' - it's not a number!", str)
	}
	if value < 1 || value > 10 {
		log.Fatalf("'%s' - input MUST be numbers from 1 to 10 inclusive!", str)
	}

	return value
}

func validateOperation(operator string) {
	if operation.Operations[operator] == nil {
		log.Fatal("Operation not found!")
	}
}
