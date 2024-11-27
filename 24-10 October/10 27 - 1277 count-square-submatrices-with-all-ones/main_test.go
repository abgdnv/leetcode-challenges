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
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1}},
			15},
		{[][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0}},
			7},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := countSquares(tC.matrix)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
