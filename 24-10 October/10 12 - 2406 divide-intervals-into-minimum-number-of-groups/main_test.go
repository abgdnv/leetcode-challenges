package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		want      int
	}{

		{[][]int{{441459, 446342}, {801308, 840640}, {871890, 963447}, {228525, 336985}, {807945, 946787}, {479815, 507766}, {693292, 944029}, {751962, 821744}}, 4},
		{[][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}, 3},
		{[][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}, 1},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := minGroups(tC.intervals)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
