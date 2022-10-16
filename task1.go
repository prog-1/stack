package main

import (
	"fmt"
)

func polish(s []string) int {
	var t []int
	for _, v := range s {
		if v == "+" {
			a := t[len(t)-2] + t[len(t)-1]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if v == "-" {
			a := t[len(t)-2] - t[len(t)-1]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if v == "*" {
			a := t[len(t)-2] * t[len(t)-1]
			t = t[:len(t)-2]
			t = append(t, a)
		} else if v == "/" {
			a := t[len(t)-2] / t[len(t)-1]
			t = t[:len(t)-2]
			t = append(t, a)
		} else {
			x := 0
			fmt.Sscan(v, &x)
			t = append(t, x)
		}
	}
	return t[0]
}

// func main() {
// 	s := []string{"8", "64", "8", "/", "*"}
// 	fmt.Println(polish(s))
// }
