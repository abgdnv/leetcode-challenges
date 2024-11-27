package main

import (
	"math"
	"slices"
)

func minimumMountainRemovals(nums []int) int {
	numsRev := make([]int, len(nums))
	copy(numsRev, nums)
	slices.Reverse(numsRev)
	lisLen := lengthOfLIS(nums)
	ldsLen := lengthOfLIS(numsRev)
	slices.Reverse(ldsLen)
	n := len(nums)
	minRemovals := math.MaxInt
	for i := range nums {
		if lisLen[i] > 1 && ldsLen[i] > 1 {
			minRemovals = min(minRemovals, n-lisLen[i]-ldsLen[i]+1)
		}
	}
	return minRemovals
}

func lengthOfLIS(nums []int) []int {
	lisLen := make([]int, len(nums))
	stack := make([]int, 0, len(nums))
	for i, num := range nums {
		if len(stack) == 0 || num > stack[len(stack)-1] {
			stack = append(stack, num)
		} else {
			idx, _ := slices.BinarySearch(stack, num)
			stack[idx] = num
		}
		lisLen[i] = len(stack)
	}
	return lisLen
}
