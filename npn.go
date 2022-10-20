package main

import (
	"fmt"
	"strconv"
)

func mainnpn() {
	fmt.Println(npn([]string{"/", "*", "+", "3", "4", "7", "49"}))
}

func npn(expr []string) int {
	res, _ := pn(0, 0, expr)
	return res
}

func pn(n, i int, expr []string) (int, int) {
	switch expr[i] {
	case "+":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 + n2, i2
	case "-":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 - n2, i2
	case "*":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 * n2, i2
	case "/":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 / n2, i2
	default:
		n, _ := strconv.Atoi(expr[i])
		return n, i
	}

}
