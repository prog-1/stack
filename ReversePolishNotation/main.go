package main

import "strconv"

func rpn(expr []string) int {
	var stack []int
	var res int
	i := -1
	for _, v := range expr {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			if v == "+" {
				res = stack[i-1] + stack[i]
			} else if v == "-" {
				res = stack[i-1] - stack[i]
			} else if v == "*" {
				res = stack[i-1] * stack[i]
			} else if v == "/" {
				res = stack[i-1] / stack[i]
			}
			stack = append(stack[:i-1], stack[i+1:]...)
			stack = append(stack, res)
			i--
		} else {
			tmp, _ := strconv.Atoi(v)
			stack = append(stack, tmp)
			i++
		}
	}
	return res
}
