package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"|(&(t,f,t),!(t))", false},
		{"|(&(t,f,t),t)", true},
		{"&(|(f))", false},
		{"|(f,f,f,t)", true},
		{"!(&(f,t))", true},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := parseBoolExpr(tC.s)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
