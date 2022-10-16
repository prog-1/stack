package main

import (
	"fmt"
	"testing"
)

func TestNpn(t *testing.T) {
	for _, tc := range []struct {
		input []string
		want  int
	}{
		{[]string{"1"}, 1},
		{[]string{"2"}, 2},
		{[]string{"123"}, 123},
		{[]string{"+", "2", "*", "2", "2"}, 6},
		{[]string{"*", "2", "+", "2", "2"}, 8},
		{[]string{"+", "1", "+", "2", "3"}, 6},
		{[]string{"+", "4", "/", "11", "3"}, 7},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := npn(tc.input); got != tc.want {
				t.Errorf("npn(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
