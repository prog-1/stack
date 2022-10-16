package main

import (
	"fmt"
	"testing"
)

func TestRpn(t *testing.T) {
	for _, tc := range []struct {
		input []string
		want  int
	}{
		{[]string{"1"}, 1},
		{[]string{"2"}, 2},
		{[]string{"123"}, 123},
		{[]string{"1", "3", "4", "+", "+"}, 8},
		{[]string{"2", "4", "3", "-", "*"}, 2},
		{[]string{"3", "4", "+", "2", "*"}, 14},
		{[]string{"4", "11", "3", "/", "+"}, 7},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if got := rpn(tc.input); got != tc.want {
				t.Errorf("rpn(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
