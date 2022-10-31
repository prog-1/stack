package main

import "testing"

func TestNPN(t *testing.T) {
	for _, tc := range []struct {
		expr []string
		want int
	}{
		{[]string{"123"}, 123},
		{[]string{"*", "+", "3", "4", "2"}, 14},
		{[]string{"*", "2", "+", "3", "4"}, 14},
		{[]string{"+", "4", "/", "11", "3"}, 7},
	} {
		t.Run(func(t *testing.T) {
			got := npn(tc.expr)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
