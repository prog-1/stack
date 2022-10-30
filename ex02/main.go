package main

import (
	"fmt"
)

func makeroot(root []string, tmp string) ([]string, string) {
	switch tmp {
	case ".":
	case "..":
		if len(root) > 0 {
			root = root[:len(root)-1]
		}
	default:
		root = append(root, tmp)

	}
	return root, ""
}
func connect(s []string) (res string) {
	if len(s) == 0 {
		return "/"
	}
	for _, v := range s {
		fmt.Println(res)
		res += "/"
		fmt.Println(res)
		res += v
	}
	return res
}
func path(s string) string {
	var root []string
	var tmp string
	for _, v := range s {
		if v == '/' {
			root, tmp = makeroot(root, tmp)
			continue
		}
		tmp += string(v)
	}
	if len(root) > 0 {
		root = root[1:]
	}
	return connect(root)

}
func main() {

}
