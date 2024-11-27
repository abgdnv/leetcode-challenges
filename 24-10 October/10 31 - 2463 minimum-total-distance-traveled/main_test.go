package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		robot   []int
		factory [][]int
		want    int64
	}{
		{[]int{9, 11, 99, 101}, [][]int{{10, 1}, {7, 1}, {14, 1}, {100, 1}, {96, 1}, {103, 1}}, 6},
		{[]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}, 4},
		{[]int{1, -1}, [][]int{{-2, 1}, {2, 1}}, 2},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := minimumTotalDistance(tC.robot, tC.factory)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
