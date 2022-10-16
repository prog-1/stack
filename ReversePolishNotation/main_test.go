package main

import (
	"testing"
)

func TestRPN(t *testing.T) {
	for _, tc := range []struct {
		name string
		expr []string
		want int
	}{
		{"case-1", []string{"3", "4", "+", "2", "*"}, 14},
		{"case-2", []string{"4", "11", "3", "/", "+"}, 7},
		{"case-3", []string{"4", "11", "3", "/", "+", "5", "-"}, 2},
		{"case-4", []string{"3", "4", "+", "2", "3", "*", "-"}, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := rpn(tc.expr)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
