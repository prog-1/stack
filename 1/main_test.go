package main

import "testing"

func TestRpn(t *testing.T) {
	for _, tc := range []struct {
		s    []string
		want int
	}{
		{[]string{"3", "4", "+", "2", "*"}, 14},
		{[]string{"4", "11", "3", "/", "+"}, 7},
		{[]string{"1", "2", "3", "+", "*"}, 5},
		{[]string{"100", "12", "2", "-", "/"}, 10},
		{[]string{"2", "3", "11", "5", "27", "+", "*", "/", "-"}, 2},
	} {
		if got := rpn(tc.s); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
