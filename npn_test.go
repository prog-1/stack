package main

import (
	"strconv"
	"testing"
)

func npn(expr []string) int {
	res, _ := pn(0, 0, expr)
	return res
}

func pn(n, i int, expr []string) (int, int) {
	switch expr[i] {
	case "+":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 + n2, i2
	case "-":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 - n2, i2
	case "*":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 * n2, i2
	case "/":
		n1, i1 := pn(n, i+1, expr)
		n2, i2 := pn(n, i1+1, expr)
		return n1 / n2, i2
	default:
		n, _ := strconv.Atoi(expr[i])
		return n, i
	}

}

func TestNPN(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []string
		want  int
	}{
		{"1", []string{"*", "+", "3", "4", "2"}, 14},
		{"1", []string{"*", "2", "+", "3", "4"}, 14},
		{"1", []string{"+", "4", "/", "11", "3"}, 7},
		{"1", []string{"+", "-4", "*", "-2", "30"}, -64},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := npn(tc.input)
			if got != tc.want {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}

func main() {
	testing.Main(
		/* matchString */ func(a, b string) (bool, error) { return a == b, nil },
		/* tests */ []testing.InternalTest{
			{Name: "Test Normal Polish Notation", F: TestNPN},
		},
		/* benchmarks */ nil /* examples */, nil)
}
