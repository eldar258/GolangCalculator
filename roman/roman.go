package roman

var ones = [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var tens = [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var hunds = [2]string{"", "C"}

var NumOneToTen = map[string]int{
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

func ToRoman(num int) string {
	h := num / 100 % 10
	t := num / 10 % 10
	o := num % 10

	return hunds[h] + tens[t] + ones[o]
}

func IsRomanNumOneToTen(num string) bool {
	return NumOneToTen[num] != 0
}
