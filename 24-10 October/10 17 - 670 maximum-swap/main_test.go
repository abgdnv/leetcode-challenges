package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		num  int
		want int
	}{
		{2736, 7236},
		{98368, 98863},
		{9973, 9973},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := maximumSwap(tC.num)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
