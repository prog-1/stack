package main

import (
	"fmt"
	"strconv"
)

func checkfornum(s string) bool {
	switch s {
	case "+":
		return false
	case "-":
		return false
	case "*":
		return false
	case "/":
		return false
	}
	return true
}

func a(s []string, i, j int) []string {
	s[i] = fmt.Sprint(j)
	s = append(s[:i+1], s[i+3:]...)
	return s
}

func npn(expr []string) int {
	for i, j := range expr {
		switch j {
		case "+":
			if checkfornum(expr[i+1]) && checkfornum(expr[i+2]) {
				n, _ := strconv.Atoi(expr[i+1])
				m, _ := strconv.Atoi(expr[i+2])
				npn(a(expr, i, n+m))
			}
		case "-":
			if checkfornum(expr[i+1]) && checkfornum(expr[i+2]) {
				n, _ := strconv.Atoi(expr[i+1])
				m, _ := strconv.Atoi(expr[i+2])
				npn(a(expr, i, n-m))
			}
		case "*":
			if checkfornum(expr[i+1]) && checkfornum(expr[i+2]) {
				n, _ := strconv.Atoi(expr[i+1])
				m, _ := strconv.Atoi(expr[i+2])
				npn(a(expr, i, n*m))
			}
		case "/":
			if checkfornum(expr[i+1]) && checkfornum(expr[i+2]) {
				n, _ := strconv.Atoi(expr[i+1])
				m, _ := strconv.Atoi(expr[i+2])
				npn(a(expr, i, n/m))
			}
		}
	}
	a, _ := strconv.Atoi(expr[0])
	return a
}

func main() {
	fmt.Println(npn([]string{"*", "2", "+", "3", "4"}))
	fmt.Println(npn([]string{"-", "5", "*", "6", "7"}))
}
