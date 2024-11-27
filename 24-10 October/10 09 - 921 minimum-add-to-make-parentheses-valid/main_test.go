package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{"()))((", 4},
		{"())", 1},
		{"(((", 3},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := minAddToMakeValid(tC.s)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
