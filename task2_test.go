package main

import (
	"testing"
)

func TestPath(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"/foo/./bar/../../baz/", "/baz"},
		{"/../", "/"},
		{"/home/", "/home"},
		{"/foo", "/foo"},
		{"/foo/./bar/baz/boo/../../../../../../", "/"},
	} {
		if got := path(tc.s); got != tc.want {
			t.Errorf("path(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
