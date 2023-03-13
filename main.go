package main

import (
	"bufio"
	"calculator/calculator"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	result := calculator.Calculate(scanner.Text())

	println(result)
}
