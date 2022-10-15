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

func rpn(expr []string) int {
	var stack []int
	for _, i := range expr {
		if opr, ok := operators[i]; !ok {
			x, _ := strconv.Atoi(i)
			stack = append(stack, x)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = opr(a, b)
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(rpn([]string{"1", "3", "4", "+", "+"}))
	fmt.Println(rpn([]string{"3", "4", "+", "2", "*"}))
}
