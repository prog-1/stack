package main

import "testing"

func TestNPN(t *testing.T) {
	for _, tc := range []struct {
		name string
		expr []string
		want int
	}{
		{"case-1", []string{"123"}, 123},
		{"case-2", []string{"*", "+", "3", "4", "2"}, 14},
		{"case-3", []string{"*", "2", "+", "3", "4"}, 14},
		{"case-4", []string{"+", "4", "/", "11", "3"}, 7},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := npn(tc.expr)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
