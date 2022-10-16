package main

import (
	"testing"
)

func TestSimplepath(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"/home/", "/home"},
		{"/foo/./bar/../../baz/", "/baz"},
	} {
		if got := simplepath(tc.s); got != tc.want {
			t.Errorf("path(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
