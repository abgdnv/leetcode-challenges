package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{[]int{6, 0, 8, 2, 1, 5}, 4}, //(i, j) = (1, 5)
		{[]int{1, 0}, 0},
		{[]int{0, 1}, 1},
		{[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7}, //(i, j) = (2, 9)
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := maxWidthRamp(tC.nums)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
