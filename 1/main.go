package main

import (
	"fmt"
	"strconv"
)

func rpn(expr []string) (result int) {
	var a []int
	for _, s := range expr {
		if s != "/" && s != "*" && s != "+" && s != "-" {
			res, _ := strconv.Atoi(s)
			a = append(a, res)
		} else if s == "/" {
			result = a[len(a)-2] / a[len(a)-1]
			a = a[:len(a)-2]
			a = append(a, result)
		} else if s == "*" {
			result = a[len(a)-2] * a[len(a)-1]
			a = a[:len(a)-2]
			a = append(a, result)
		} else if s == "+" {
			result = a[len(a)-2] + a[len(a)-1]
			a = a[:len(a)-2]
			a = append(a, result)
		} else if s == "-" {
			result = a[len(a)-2] - a[len(a)-1]
			a = a[:len(a)-2]
			a = append(a, result)
		}
	}
	return result
}

func main() {
	expr := []string{"4", "11", "3", "/", "+"}
	fmt.Println(rpn(expr))
}
