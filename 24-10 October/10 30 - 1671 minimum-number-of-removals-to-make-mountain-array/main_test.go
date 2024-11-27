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
		{[]int{4, 3, 2, 1, 1, 2, 3, 1}, 4},
		{[]int{1, 3, 1}, 0},
		{[]int{2, 1, 1, 5, 6, 2, 3, 1}, 3},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := minimumMountainRemovals(tC.nums)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
