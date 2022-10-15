package main

import (
	"fmt"
	"strings"
)

// Simplifies path
func simplify(path string) string {
	var dirs []string
	path = strings.Trim(path, "/")
	for _, el := range strings.Split(path, "/") {
		switch el {
		case ".": // Breaks from switch
		case "":
		case "..":
			if len(dirs) > 0 {
				dirs = dirs[:len(dirs)-1] // Pop
			}
		default: // Dir name
			dirs = append(dirs, el)
		}
	}

	var result string
	for _, dir := range dirs {
		result += "/" + dir
	}
	if result == "" {
		return "/"
	}
	return result
}

func main() {
	path := "/../"
	fmt.Println(simplify(path))
}
