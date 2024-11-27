package main

/**
 * Definition for a binary tree node.
 **/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {

	//count sum for each level
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

	root.Val = 0
	queue = queue[:1]
	queue[0] = root
	level := 0
	for len(queue) > 0 {
		level++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			leftVal := 0
			rightVal := 0
			if node.Left != nil {
				leftVal = node.Left.Val
			}
			if node.Right != nil {
				rightVal = node.Right.Val
			}
			updateValue := 0
			if node.Left != nil || node.Right != nil {
				updateValue = int(sum[level] - int64(leftVal+rightVal))
			}

			if node.Left != nil {
				node.Left.Val = updateValue
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = updateValue
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
