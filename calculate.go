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

func main() {
	var x, plus, y string
	fmt.Scanln(&x, &plus, &y)
	num1, num2 := parseInt(x, y)
	if plus == "+" {
		fmt.Println(add(num1, num2))
	}
}
