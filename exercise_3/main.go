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

func doOperations(expr []string) (newExpr []string, result int) {
	// if len(expr) == 0 {   we don't need this if-statement because we assume that input is always correct
	//	   return
	// }
	if opr, ok := operators[expr[0]]; ok {
		newExpr, a := doOperations(expr[1:]) // first operand
		_, b := doOperations(newExpr[1:])    // second operand
		return newExpr[1:], opr(a, b)
	} else {
		result, _ = strconv.Atoi(expr[0])
		return expr, result
	}
}

func npn(expr []string) int {
	_, result := doOperations(expr)
	return result
}

func main() {
	fmt.Println(npn([]string{"123"}))
	fmt.Println(npn([]string{"*", "+", "3", "4", "2"}))
	fmt.Println(npn([]string{"+", "*", "+", "3", "4", "2", "3"}))
	fmt.Println(npn([]string{"+", "2", "*", "2", "2"}))
}
