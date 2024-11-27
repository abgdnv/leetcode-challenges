package main

import (
	"slices"
)

func minGroups(intervals [][]int) int {
	if len(intervals) == 1 {
		return 1
	}
	points := make([][]int, 0)
	for _, v := range intervals {
		points = append(points, []int{v[0], 1})
		points = append(points, []int{v[1] + 1, -1})

	}
	slices.SortFunc(points, func(i, j []int) int {
		if i[0] == j[0] {
			return i[1] - j[1]
		}
		return i[0] - j[0]
	})
	groups := 0
	balance := 0
	for i := range len(points) {
		balance += points[i][1]
		groups = max(groups, balance)
	}
	return groups
}

// bruteforce
// func minGroups(intervals [][]int) int {
// 	slices.SortFunc(intervals, func(i, j []int) int {
// 		if i[0] == j[0] {
// 			return i[1] - j[1]
// 		}
// 		return i[0] - j[0]
// 	})
// 	groups := make([][][]int, 0)
// 	for _, v := range intervals {
// 		var intersection = false
// 		for j := range groups {
// 			intersection = false
// 			for k := range groups[j] {
// 				if v[0] >= groups[j][k][0] && v[0] <= groups[j][k][1] || v[1] <= groups[j][k][0] && v[1] >= groups[j][k][1] {
// 					intersection = true
// 					break
// 				}
// 			}
// 			if !intersection {
// 				groups[j] = append(groups[j], v)
// 				break
// 			}
// 		}
// 		if intersection || len(groups) == 0 {
// 			groups = append(groups, [][]int{v})
// 		}
// 	}
// 	return len(groups)
// }
