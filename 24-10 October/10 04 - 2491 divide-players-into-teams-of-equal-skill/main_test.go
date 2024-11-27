package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		want int64
	}{
		{[]int{3, 2, 5, 1, 3, 4}, 22},
		{[]int{3, 4}, 12},
		{[]int{1, 1, 2, 3}, -1},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.arr, " - ", tC.want), func(t *testing.T) {

		})
	}
}
