package main

import "testing"

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    string
		want string
	}{
		{"case-1", "/home/", "/home"},
		{"case-2", "/../", "/"},
		{"case-3", "/foo/./bar/../../baz/", "/baz"},
		{"case-4", "///foo/./bar/../../baz/", "/baz"},
		{"case-5", "./foo/./bar//./../baz//.", "/foo/baz"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := simplify(tc.s)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
