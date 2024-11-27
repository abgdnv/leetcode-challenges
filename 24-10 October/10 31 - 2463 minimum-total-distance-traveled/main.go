package main

import (
	"math"
	"slices"
)

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	robotsLen := len(robot)
	slices.Sort(robot)
	slices.SortFunc(factory, func(a, b []int) int {
		return a[0] - b[0]
	})
	fexLen := 0
	for i := range factory {
		fexLen += factory[i][1]
	}
	fex := make([]int, 0, fexLen)
	for i := range factory {
		for range factory[i][1] {
			fex = append(fex, factory[i][0])
		}
	}

	next := make([]int64, len(fex)+1)
	current := make([]int64, len(fex)+1)
	current[fexLen] = 1e12

	for i := robotsLen - 1; i >= 0; i-- {
		if i != robotsLen-1 {
			next[fexLen] = 1e12
		}
		for j := fexLen - 1; j >= 0; j-- {
			assign := int64(math.Abs(float64(robot[i]-fex[j]))) + next[j+1]
			skip := current[j+1]
			current[j] = min(assign, skip)
		}
		copy(next, current)
	}
	return current[0]
}
