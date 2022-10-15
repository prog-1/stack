package main

import (
	"fmt"
	"strconv"
)

var operators = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func doOpeartions(expr []string, index int, currentResult int) int {
	if index <= 0 {
		return currentResult
	}
	opr := operators[expr[index-1]]
	x, _ := strconv.Atoi(expr[index])
	return doOpeartions(expr, index-2, opr(x, currentResult))
}

func npn(expr []string) (result int) {
	currentResult, _ := strconv.Atoi(expr[len(expr)-1])
	return doOpeartions(expr, len(expr)-2, currentResult)
}

func main() {
	fmt.Println(npn([]string{"*", "2", "+", "2", "2"}))
	fmt.Println(npn([]string{"+", "1", "+", "2", "3"}))
}
