package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{40, 10, 20, 30},
			[]int{4, 1, 2, 3},
		},
		{
			[]int{100, 100, 100},
			[]int{1, 1, 1},
		},
		{
			[]int{37, 12, 28, 9, 100, 56, 80, 5, 12},
			[]int{5, 3, 4, 2, 8, 6, 7, 1, 3},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := arrayRankTransform(tC.arr)
			if !isEqual(value, tC.want) {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}

func isEqual(s1, s2 []int) bool {
	return reflect.DeepEqual(s1, s2)
}
