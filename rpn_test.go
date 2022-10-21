package main

import (
	"strconv"
	"testing"
)

func rpn(expr []string) int {
	var stack []int
	for _, el := range expr {
		switch el {
		case "+":
			res := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "-":
			res := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "*":
			res := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case "/":
			res := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default: // number
			n, _ := strconv.Atoi(el)
			stack = append(stack, n)
		}
	}
	return stack[0]
}

func TestRPN(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []string
		want  int
	}{
		{"1", []string{"4", "3", "3", "*", "*"}, 36},
		{"2", []string{"4", "11", "3", "/", "+"}, 7},
		{"3", []string{"3", "4", "+", "2", "*"}, 14},
		{"4", []string{"2", "4", "3", "+", "*", "5", "-"}, 9},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := rpn(tc.input)
			if got != tc.want {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}

func mainRPN() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Reverse Polish Notation", F: TestRPN},
		},
		/* benchmarks */ nil /* examples */, nil)
}
