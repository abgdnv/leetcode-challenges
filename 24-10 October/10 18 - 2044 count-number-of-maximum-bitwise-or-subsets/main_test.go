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
		{[]int{3, 1}, 2},
		{[]int{2, 2, 2}, 7},
		{[]int{3, 2, 1, 5}, 6},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := countMaxOrSubsets(tC.nums)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
