package main

import (
	"math"
	"slices"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	idxs := make([]int, len(arr))
	result := make([]int, len(arr))

	for i := range len(arr) {
		idxs[i] = i
	}

	slices.SortFunc(idxs, func(i, j int) int {
		return arr[i] - arr[j]
	})
	slices.SortFunc(arr, func(i, j int) int {
		return i - j
	})

	rating := 0
	prevValue := math.MinInt
	for i := range len(arr) {
		if arr[i] > prevValue {
			rating++
			prevValue = arr[i]
		}
		arr[i] = rating
	}

	for i := range arr {
		result[idxs[i]] = arr[i]
	}

	return result
}
