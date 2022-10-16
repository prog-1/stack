package main

import (
	"strings"
)

func path(s string) string {
	var tmp []string
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == ".." || string(s[i]) == "." {
			break
		}
		tmp = append(tmp, string(s[i]))

	}
	if tmp[0] == "/" && len(tmp) > 1 {
		tmp = append(tmp[:0], tmp[1:]...)
	}
	if len(tmp) > 1 {
		for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		}
	}
	s = strings.Join(tmp, "")
	return s
}

// func main() {
// 	s := "/../"
// 	fmt.Println(path(s))
// }
