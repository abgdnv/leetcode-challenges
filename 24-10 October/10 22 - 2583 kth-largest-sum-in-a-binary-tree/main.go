package main

import "slices"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := make([]*TreeNode, 1, 1000000)
	queue[0] = root
	sum := []int64{}
	var levelSum int64
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum = 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			levelSum += int64(node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		sum = append(sum, levelSum)
	}
	if k <= len(sum) {
		slices.SortFunc(sum, func(a, b int64) int {
			return int(b - a)
		})
		return sum[k-1]
	}
	return -1
}

// func kthLargestLevelSum(root *TreeNode, k int) int64 {
// 	sum = []int64{}
// 	countLevelSum(root, 0)
// 	if k <= len(sum) {
// 		slices.SortFunc(sum, func(a, b int64) int {
// 			return int(b - a)
// 		})
// 		return sum[k-1]
// 	}
// 	return -1
// }

// var sum []int64

// func countLevelSum(root *TreeNode, level int) int64 {
// 	if root == nil {
// 		return 0
// 	}
// 	if level > len(sum)-1 {
// 		sum = append(sum, 0)
// 	}
// 	sum[level] += int64(root.Val)
// 	return countLevelSum(root.Left, level+1) + countLevelSum(root.Right, level+1)
// }
