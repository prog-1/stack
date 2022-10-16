package main

import "testing"

func TestRpn(t *testing.T) {
	for _, tc := range []struct {
		n    []string
		want int
	}{
		{[]string{"4", "11", "3", "/", "+"}, 7},
		{[]string{"3", "4", "+", "2", "*"}, 14},
		{[]string{"3", "3", "3", "+", "*"}, 18},
		{[]string{"3", "3", "3", "*", "*"}, 27},
		{[]string{"3", "3", "3", "/", "*"}, 3},
	} {
		if got := rpn(tc.n); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
