package operation

var Operations = map[string]func(a int, b int) int{
	"+": sum,
	"/": div,
	"*": mult,
	"-": sub,
}

func sum(a int, b int) int {
	return a + b
}

func div(a int, b int) int {
	return a / b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}
