package main

import "fmt"

func simplifyAbsolutePath(path string) (result string) {
	var stack []string
	var dir string
	for index, i := range path {
		if i != '/' && i != '.' {
			dir += string(i)
		} else if i == '.' && path[index+1] == '.' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if dir != "" {
			stack = append(stack, dir)
			dir = ""
		}
	}
	for _, i := range stack {
		result += "/" + i
	}
	if len(result) == 0 {
		return "/"
	}
	return
}

func main() {
	fmt.Println(simplifyAbsolutePath("/home/"))
	fmt.Println(simplifyAbsolutePath("/foo/./bar/../../baz/"))
}
