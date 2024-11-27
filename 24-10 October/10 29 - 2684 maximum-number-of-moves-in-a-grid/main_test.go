package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{
			{2, 4, 3, 5},
			{5, 4, 9, 3},
			{3, 4, 2, 11},
			{10, 9, 13, 15}}, 3},
		{[][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}, 0},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := maxMoves(tC.matrix)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
