package main

import (
	"strings"
)

func simplify(s string) string {
	var res string
	path := strings.Split(s, "/")
	for i := 0; i < len(path); i++ {
		if path[i] == "" || path[i] == "." {
			path = append(path[:i], path[i+1:]...)
			i--
		} else if path[i] == ".." {
			if i > 0 {
				path = append(path[:i-1], path[i+1:]...)
				i -= 2
			} else {
				path = append(path[:i], path[i+1:]...)
				i--
			}
		}
	}
	for _, v := range path {
		res += "/" + v
	}
	if res == "" {
		return "/"
	}
	return res
}
