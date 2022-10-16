package main

import (
	"testing"
)

func TestNPN(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []string
		want  int
	}{
		{"nil", nil, 0},
		{"empty", []string{}, 0},
		{"1", []string{"123"}, 123},
		{"2", []string{"*", "+", "3", "4", "2"}, 14},
		{"3", []string{"+", "4", "/", "11", "3"}, 7},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := npn(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
