package main

import "testing"

func TestNpn(t *testing.T) {
	for _, tc := range []struct {
		n    []string
		want int
	}{

		{[]string{"*", "+", "3", "4", "2"}, 14},
		{[]string{"*", "2", "+", "3", "4"}, 14},
		{[]string{"+", "4", "/", "11", "3"}, 7},
		{[]string{"+", "100", "/", "99", "3"}, 133},
		{[]string{"+", "1", "+", "2", "3"}, 6},
		{[]string{"*", "25", "*", "100", "-1"}, -2500},
	} {
		if got := npn(tc.n); got != tc.want {
			t.Errorf("npn(%v) = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
