package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		matrix []int
		want   int
	}{
		{[]int{4, 3, 6, 16, 8, 2}, 3},
		{[]int{2, 3, 5, 6, 7}, -1},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := longestSquareStreak(tC.matrix)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
