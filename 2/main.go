package main

import (
	"fmt"
)

func getres(p string) (res string) {
	var s []string
	s = append(s, p)
	for _, s2 := range s {
		res += "/" + s2
	}
	return res
}

func simplepath(s string) (result string) {
	var stack []string
	var path string
	for i := 0; i < len(s); i++ {
		if s[i] == '.' && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		} else if s[i] != '/' && s[i] != '.' {
			path += string(s[i])
		} else if path != "" {
			result = getres(path)
			path = ""
		}
	}
	return result
}

func main() {
	s := "/foo/./bar/../../baz/"
	fmt.Println(simplepath(s))
}
