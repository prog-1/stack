package main

import "testing"

func TestRpn(t *testing.T) {
	for _, tc := range []struct {
		n    []string
		want int
	}{

		{[]string{"2", "2", "2", "*", "+"}, 6},
		{[]string{"3", "2", "2", "*", "+"}, 7},
		{[]string{"2", "2", "2", "+", "*"}, 8},
		{[]string{"3", "4", "-"}, -1},
		{[]string{"2", "2", "2", "/", "+"}, 3},
		{[]string{"6", "5", "4", "*", "*"}, 120},
	} {
		if got := rpn(tc.n); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
