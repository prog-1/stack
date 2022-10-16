package main

import "testing"

func TestPolish(t *testing.T) {
	for _, tc := range []struct {
		n    []string
		want int
	}{

		{[]string{"4", "11", "3", "/", "+"}, 7},
		{[]string{"3", "4", "+", "2", "*"}, 14},
		{[]string{"7", "2", "/", "-3", "*"}, -9},
		{[]string{"100", "101", "-"}, -1},
		{[]string{"8", "64", "8", "/", "*"}, 64},
		{[]string{"25", "100", "-1", "/", "*"}, -2500},
	} {
		if got := polish(tc.n); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
