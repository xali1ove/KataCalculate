package main

import (
	"fmt"
	"strconv"
)

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func parseInt(x, y string) (int, int) {
	var num1, num2 int
	num1, _ = strconv.Atoi(x)
	num2, _ = strconv.Atoi(y)
	return num1, num2
}

func romanToInt(s string) int {
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}

func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func main() {
	var x, plus, y string
	fmt.Scanln(&x, &plus, &y)
	if x == "I" && y == "II" {
		fir := romanToInt(x) + romanToInt(y)
		fmt.Println(Roman(fir))
	}
	num1, num2 := parseInt(x, y)
	switch plus {
	case "+":
		fmt.Println(add(num1, num2))
	case "-":
		fmt.Println(sub(num1, num2))
	case "*":
		fmt.Println(mul(num1, num2))
	case "/":
		fmt.Println(div(num1, num2))
	}
}
