package main

import (
	"fmt"
	"strconv"
)

func eval(expr []string) ([]string, int) {
	if expr[0] == "/" || expr[0] == "*" || expr[0] == "+" || expr[0] == "-" {
		expr2, a := eval(expr[1:])
		expr3, b := eval(expr2)
		if expr[0] == "/" {
			return expr3, a / b
		} else if expr[0] == "*" {
			return expr3, a * b
		} else if expr[0] == "+" {
			return expr3, a + b
		} else if expr[0] == "-" {
			return expr3, a - b
		}
	}
	x, _ := strconv.Atoi(expr[0])
	return expr[1:], x
}

func npn(expr []string) (result int) {
	_, result = eval(expr)
	return result
}

func main() {
	expr := []string{"*", "+", "3", "4", "2"}
	fmt.Println(npn(expr))
}
