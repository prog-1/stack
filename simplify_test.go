package main

import (
	"reflect"
	"strings"
	"testing"
)

func simplify(str string) (res string) {
	var stack []string
	str = strings.Trim(str, "/")    //get rid of '/' in edges
	path := strings.Split(str, "/") //split all dirs from string
	for _, dir := range path {
		if dir == "." || dir == "" {
			// I ain't worried about it to-do-do-do-too
		} else if dir == ".." && len(stack) != 0 {
			stack = stack[:len(stack)-1]
		} else if dir != ".." {
			stack = append(stack, dir)
		}
	}

	for _, el := range stack {
		res += "/" + el
	}
	if res == "" {
		res = "/"
	}
	return res
}

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"/hed/./ge/../../hog/", "/hog"},
		{"/foo/./bar/../../baz/", "/baz"},
		{"/work/./done/", "/work/done"},
		{"/home/", "/home"},
		{"", "/"},
		{"/", "/"},
		{"/../", "/"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			got := simplify(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Simplify", F: TestSimplify},
		},
		/* benchmarks */ nil /* examples */, nil)
}
