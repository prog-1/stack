package main

import "testing"

func TestRpn(t *testing.T) {
	for _, tc := range []struct {
		input []string
		want  int
	}{
		{[]string{"3", "4", "+", "2", "*"}, 14},
		{[]string{"4", "11", "3", "/", "+"}, 7},
		{[]string{"4", "11", "3", "/", "-"}, 1},
		{[]string{"4", "11", "3", "/", "*"}, 12},
	} {
		if got := rpn(tc.input); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.input, got, tc.want)
		}
	}
}

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"/home/./richard/./Desktop/../Downloads", "/home/richard/Downloads"},
		{"/home/./richard/.//Desktop/../Downloads", "/home/richard/Downloads"},
		{"/../Downloads//articles", "/Downloads/articles"},
	} {
		if got := simplify(tc.input); got != tc.want {
			t.Errorf("rpn(%v) = %v, want = %v", tc.input, got, tc.want)
		}
	}
}
