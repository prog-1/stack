package main

import (
	"fmt"
	"testing"
)

func TestSimplifyAbsolutePath(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"/", "/"},
		{"/home/", "/home"},
		{"/../", "/"},
		{"/foo/./bar/../../baz/", "/baz"},
		{"/foo/./bar/", "/foo/bar"},
		{"//foo//.///bar//", "/foo/bar"}, // multiple consecutive slashes are treated as a single slash
		{"/foo/../", "/"},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := simplifyAbsolutePath(tc.input); got != tc.want {
				t.Errorf("rpn(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
