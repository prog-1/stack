package main

import "fmt"

func add(s []string, path string, i int) ([]string, int) {
	var cnt int
	var a string
	for path[i] != '/' {
		a += string(path[i])
		i++
		cnt++
		if i == len(path) {
			break
		}
	}
	s = append(s, a)
	return s, cnt - 1
}

func simplepath(path string) string {
	var cnt int
	var out string
	var s []string
	for i := 0; i <= len(path)-1; i++ {
		switch path[i] {
		case '.':
			if path[i+1] == '.' && len(s) >= 1 {
				i++
				s = s[:len(s)-1]
			}
			continue
		case '/':
			continue
		}
		s, cnt = add(s, path, i)
		i += cnt
	}

	if len(s) < 1 {
		return out + "/"
	}
	for _, j := range s {
		out += "/" + j
	}
	return out
}

func main() {
	fmt.Println(simplepath("/foo/./bar/../../baz/boo/"))
}
