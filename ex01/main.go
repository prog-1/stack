package main

import (
	"strconv"
)

func rpn(s []string) int {
	var a []int
	for _, v := range s {
		var res int
		if v != "/" && v != "*" && v != "+" && v != "-" {
			res, _ := strconv.Atoi(v)
			a = append(a, res)
		} else {
			switch v {
			case "+":
				res = a[len(a)-2] + a[len(a)-1]
				a = a[:len(a)-2]
				a = append(a, res)

			case "-":
				res = a[len(a)-2] - a[len(a)-1]
				a = a[:len(a)-2]
				a = append(a, res)
			case "/":
				res = a[len(a)-2] / a[len(a)-1]
				a = a[:len(a)-2]
				a = append(a, res)
			case "*":
				res = a[len(a)-2] * a[len(a)-1]
				a = a[:len(a)-2]
				a = append(a, res)
			}
		}
	}
	return a[0]
}
