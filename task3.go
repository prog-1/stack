package main

import (
	"fmt"
)

func npn(n []string) int {
	var t []int
	for v := len(n) - 1; v >= 0; v-- {
		if n[v] == "+" {
			a := t[len(t)-1] + t[len(t)-2]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if n[v] == "-" {
			a := t[len(t)-1] - t[len(t)-2]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if n[v] == "*" {
			a := t[len(t)-1] * t[len(t)-2]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if n[v] == "/" {
			a := t[len(t)-1] / t[len(t)-2]
			t = t[:len(t)-2]
			t = append(t, a)
		} else {
			x := 0
			fmt.Sscan(n[v], &x)
			t = append(t, x)
		}
	}
	return t[0]
}

func main() {
	s := []string{"+", "100", "/", "99", "3"}
	fmt.Println(npn(s))
}
