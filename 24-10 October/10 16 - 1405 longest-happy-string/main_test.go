package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		a, b, c int
		want    string
	}{
		{0, 8, 11, "ccbccbbccbbccbbccbc"},
		{1, 1, 7, "ccaccbcc"},
		{7, 1, 0, "aabaa"},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := longestDiverseString(tC.a, tC.b, tC.c)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
