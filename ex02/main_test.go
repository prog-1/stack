package main

import "testing"

func TestPath(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/foo/./bar/../../baz/", "/baz"},
		{"/home/./foo/../../baz/bar/", "/baz/bar"},
		{"/home/bar/baz/boo/../../../../", "/"},
		{"/home/./foo/../../bar/baz/./././../", "/bar"},
	} {
		if got := simplify(tc.n); got != tc.want {
			t.Errorf("path(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
