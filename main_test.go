package main

import "testing"

func TestChecParenthesis(t *testing.T) {
	var tests = []struct {
		expression string
		want       bool
	}{
		{"1*2+(5*3)-(2+(3*5))", true},
		{"1*2+(5*3)-(2+(3*5)", false},
		{"1*2+(5*3)-(2+(3*5))(", false},
		{"1*2+(5*3)-{2+)3*5))()", false},
		{"1*2+(5*3)-(2+(3*5))[}", false},
		{"1*2+(5*3)-(2+(3*5))[]", true},
	}
	for _, test := range tests {
		// sub test
		t.Run(test.expression, func(t *testing.T) {
			if got := checkParenthesis(test.expression); got != test.want {
				t.Errorf("checkParenthesis(%q) = %v", test.expression, got)
			}
		})
	}
}
