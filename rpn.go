package main

import (
	"fmt"
	"strconv"
)

func mainrpn() {
	fmt.Println(rpn([]string{"3", "4", "+", "7", "*"}))
}

func rpn(expr []string) int {
	var stack []int
	for _, i := range expr {
		switch i {
		case "+":
			stack = append(stack, stack[0]+stack[1])
			stack = stack[2:]
		case "-":
			stack = append(stack, stack[0]-stack[1])
			stack = stack[2:]
		case "*":
			stack = append(stack, stack[0]*stack[1])
			stack = stack[2:]
		case "/":
			stack = append(stack, stack[0]/stack[1])
			stack = stack[2:]
		default: // number
			n, _ := strconv.Atoi(i)
			stack = append(stack, n)
		}

	}
	return stack[0]
}
