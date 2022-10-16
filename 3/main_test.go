package main

import "testing"

func TestNpn(t *testing.T) {
	for _, tc := range []struct {
		s    []string
		want int
	}{
		{[]string{"3"}, 3},
		{[]string{"*", "+", "3", "4", "2"}, 14},
		{[]string{"*", "2", "+", "3", "4"}, 14},
		{[]string{"+", "4", "/", "11", "3"}, 7},
		{[]string{"-", "/", "*", "+", "2", "3", "11", "5", "27"}, -16},
	} {
		if got := npn(tc.s); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
