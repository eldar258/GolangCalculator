package main

import (
	"bufio"
	"calculator/operation"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a, b, operator, isRoman := validate(scanner.Text())

	result := operation.Operations[operator](a, b)

	if isRoman {
		if result < 1 {
			log.Fatal("No Roman numerals less than 'I'!")
		}
		println(toRoman(result))
	} else {
		println(result)
	}
}

var ones = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hunds = [2]string{"", "C"}

func toRoman(num int) string {
	h := num / 100 % 10
	t := num / 10 % 10
	o := num % 10

	return hunds[h] + tens[t] + ones[o]
}

func validate(expression string) (int, int, string, bool) {
	values := strings.Split(expression, " ")

	if values == nil || len(values) != 3 {
		log.Fatal("Wrong data format!", len(values))
	}

	aString := values[0]
	bString := values[2]
	operator := values[1]

	validateOperation(operator)
	var isRomanNum bool
	a, b := getNumbers(aString, bString, &isRomanNum)
	return a, b, operator, isRomanNum
}

func getNumbers(aString string, bString string, isRoman *bool) (int, int) {
	isARoman := IsRomanNum(aString)
	isBRoman := IsRomanNum(bString)

	if isARoman != isBRoman {
		log.Fatal("Numbers in different formats!")
	}
	*isRoman = isARoman

	if !isARoman {
		a := parseInt(aString)
		b := parseInt(bString)
		return a, b
	}

	return RomanNum[aString], RomanNum[bString]
}

func parseInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("'%s' - it's not a number!", str)
	}
	if value < 1 || value > 10 {
		log.Fatalf("'%s' - input MUST be numbers from 1 to 10 inclusive!\", str")
	}

	return value
}

func IsRomanNum(num string) bool {
	return RomanNum[num] != 0
}

func validateOperation(operator string) {
	if operation.Operations[operator] == nil {
		log.Fatal("Operation not found!")
	}
}

var RomanNum = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}
