package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		p    int
		want int
	}{
		{[]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148, 7},
		{[]int{26, 19, 11, 14, 18, 4, 7, 1, 30, 23, 19, 8, 10, 6, 26, 3}, 26, 3}, //buggy
		{[]int{1, 2, 3}, 7, -1},
		{[]int{1, 2, 3}, 3, 0},
		{[]int{6, 3, 5, 2}, 9, 2},
		{[]int{3, 1, 4, 2}, 6, 1},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := minSubarray(tC.nums, tC.p)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
