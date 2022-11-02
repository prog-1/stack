package main

import (
	"strings"
)

func simplify(s string) string {
	a := strings.Split(s, "/")
	var path []string
	var res string
	for _, v := range a {
		switch v {
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(path) != 0 {
				path = path[:len(path)-1]
			}
		default:
			path = append(path, v)
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
func main() {

}
