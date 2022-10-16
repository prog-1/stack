package main

import (
	"fmt"
	"strconv"
)

var operators = map[string]func(a, b int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

func isnumber(s []string, i int) bool {
	if _, ok := operators[s[i+1]]; ok {
		return false
	}
	if _, ok := operators[s[i+2]]; ok {
		return false
	}
	return true
}
func npn(s []string) int {
	for i, v := range s {
		if len(s) <= 1 {
			break
		}
		if x, ok := operators[v]; ok {
			if isnumber(s, i) {
				a, _ := strconv.Atoi(s[i+1])
				b, _ := strconv.Atoi(s[i+2])
				res := strconv.Itoa(x(a, b))
				tmp := append(s[:i], res)
				if len(s) > 1 {
					s = append(tmp, s[i+3:]...)
				} else {
					res, _ := strconv.Atoi(s[0])
					return res
				}

			}

		}

	}
	if len(s) > 1 {
		return npn(s)

	}
	res, _ := strconv.Atoi(s[0])
	return res
}
func main() {
	fmt.Println(npn([]string{"+", "4", "/", "11", "3"}))
}
