package main

import (
	"fmt"
	"strconv"
)

func atoi(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func rpn(expr []string) int {
	var stack []int
	for _, v := range expr {
		switch v {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]+stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2]*stack[len(stack)-1])
		default:
			stack = append(stack, atoi(v))
		}
		fmt.Println(stack)
	}
	return stack[0]
}

func main() {
	fmt.Println(rpn([]string{"3", "4", "+", "2", "*"}))
}
