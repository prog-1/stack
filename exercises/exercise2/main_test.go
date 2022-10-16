package main

import (
	"reflect"
	"testing"
)

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"Empty", "", "/"},
		{"//", "//", "/"},
		{"Funny", "/foo/./bar/../../baz/", "/baz"},
		{"Not funny", "/../", "/"},
		{"Home", "/home/", "/home"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := simplify(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}
