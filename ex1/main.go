package main

import (
	"strconv"
)

func rpn(s []string) int {
	var stack []int

	for _, v := range s {
		var r int
		if x, err := strconv.Atoi(v); err == nil {
			stack = append(stack, x)
		} else {
			switch v {
			case "+":
				r = stack[len(stack)-2] + stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, r)

			case "-":
				r = stack[len(stack)-2] - stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, r)
			case "/":
				r = stack[len(stack)-2] / stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, r)
			case "*":
				r = stack[len(stack)-2] * stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, r)
			}
		}
	}
	return stack[0]
}
