package main

import (
	"fmt"
	"strconv"
)

// Caclulates expression that uses Normal Polish Notation
func npn(expr []string) int {
	if len(expr) == 0 || expr == nil {
		return 0
	}
	var pn func(expr []string) (int, []string)
	pn = func(expr []string) (int, []string) {
		switch expr[0] {
		case "+":
			a, ne1 := pn(expr[1:])
			b, ne2 := pn(ne1[1:])
			return a + b, ne2
		case "-":
			a, ne1 := pn(expr[1:])
			b, ne2 := pn(ne1[1:])
			return a - b, ne2
		case "*":
			a, ne1 := pn(expr[1:])
			b, ne2 := pn(ne1[1:])
			return a * b, ne2
		case "/":
			a, ne1 := pn(expr[1:])
			b, ne2 := pn(ne1[1:])
			return a / b, ne2
		default: // Number
			val, _ := strconv.Atoi(expr[0])
			return val, expr
		}
	}

	val, _ := pn(expr)
	return val
}

func main() {
	//expr := []string{"+", "4", "/", "11", "3"}
	var expr []string = nil
	fmt.Println(npn(expr))
}
