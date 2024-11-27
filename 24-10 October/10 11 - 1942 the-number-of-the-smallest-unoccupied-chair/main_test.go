package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		times        [][]int
		targetFriend int
		want         int
	}{

		{[][]int{{7, 10}, {6, 7}, {1, 3}, {2, 7}, {4, 5}}, 0, 0},
		{[][]int{{3, 10}, {1, 5}, {2, 6}}, 0, 2},
		{[][]int{{1, 4}, {2, 3}, {4, 6}}, 1, 1},
		{[][]int{{33889, 98676}, {80071, 89737}, {44118, 52565}, {52992, 84310}, {78492, 88209}, {21695, 67063}, {84622, 95452}, {98048, 98856}, {98411, 99433}, {55333, 56548}, {65375, 88566}, {55011, 62821}, {48548, 48656}, {87396, 94825}, {55273, 81868}, {75629, 91467}},
			6,
			2},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := smallestChair(tC.times, tC.targetFriend)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
