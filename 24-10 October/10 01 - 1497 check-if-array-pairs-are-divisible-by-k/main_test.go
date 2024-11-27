package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		k    int
		want bool
	}{
		{
			[]int{-1, 1, -2, 2, -3, 3, -4, 4},
			3,
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9},
			5,
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			7,
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			10,
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := canArrange(tC.arr, tC.k)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
