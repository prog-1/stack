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
		{"/foo/./bar/../../baz/boo/", "/baz/boo"},
		{"/foo/./bar/baz/boo/../../../../../../", "/"},
		{"/foo/./bar/../../baz/././././../boo/", "/boo"},
	} {
		if got := path(tc.n); got != tc.want {
			t.Errorf("path(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
