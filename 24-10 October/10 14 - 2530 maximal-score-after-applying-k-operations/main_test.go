package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{10, 10, 10, 10, 10}, 5, 50},
		{[]int{1, 10, 3, 3, 3}, 3, 17},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := maxKelements(tC.nums, tC.k)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
