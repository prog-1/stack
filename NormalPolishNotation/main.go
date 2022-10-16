package main

import (
	"strconv"
)

func npn(expr []string) int {
	return eval(expr, len(expr)-1, 0)
}

func eval(expr []string, i, res int) int {
	var tmp1, tmp2 int
	if i < 0 {
		return res
	}
	if expr[0] != "+" && expr[0] != "-" && expr[0] != "*" && expr[0] != "/" {
		res, _ = strconv.Atoi(expr[0])
		return res
	}
	if expr[i] == "+" || expr[i] == "-" || expr[i] == "*" || expr[i] == "/" {
		if expr[i+2] == "+" || expr[i+2] == "-" || expr[i+2] == "*" || expr[i+2] == "/" {
			tmp2 = res
		} else {
			tmp2, _ = strconv.Atoi(expr[i+2])
		}
		if expr[i+1] == "+" || expr[i+1] == "-" || expr[i+1] == "*" || expr[i+1] == "/" {
			tmp1 = res
			tmp2, _ = strconv.Atoi(expr[len(expr)-1])
		} else {
			tmp1, _ = strconv.Atoi(expr[i+1])
		}
		if expr[i] == "+" {
			res = tmp1 + tmp2
		} else if expr[i] == "-" {
			res = tmp1 - tmp2
		} else if expr[i] == "*" {
			res = tmp1 * tmp2
		} else if expr[i] == "/" {
			res = tmp1 / tmp2
		}
	}
	return eval(expr, i-1, res)
}
