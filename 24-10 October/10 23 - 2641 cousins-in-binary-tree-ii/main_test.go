package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		root TreeNode
		want TreeNode
	}{
		{
			TreeNode{5, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{10, nil, nil}}, &TreeNode{9, &TreeNode{7, nil, nil}, nil}},
			TreeNode{0, &TreeNode{0, &TreeNode{7, nil, nil}, &TreeNode{7, nil, nil}}, &TreeNode{0, &TreeNode{11, nil, nil}, nil}},
		},
		{
			TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := replaceValueInTree(&tC.root)
			if !isTreesEqual(value, &tC.want) {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}

func isTreesEqual(root1, root2 *TreeNode) bool {
	queue1 := make([]*TreeNode, 1, 1000000)
	queue2 := make([]*TreeNode, 1, 1000000)
	queue1[0] = root1
	queue2[0] = root2

	for len(queue1) > 0 && len(queue2) > 0 {
		levelSize1 := len(queue1)
		levelSize2 := len(queue1)
		if levelSize1 != levelSize2 {
			fmt.Println("Diff size", levelSize1, levelSize2)
			return false
		}
		for i := 0; i < levelSize1; i++ {
			node1 := queue1[0]
			node2 := queue2[0]
			queue1 = queue1[1:]
			queue2 = queue2[1:]
			if node1.Val != node2.Val {
				fmt.Println("Diff values", node1.Val, node2.Val)
				return false
			}
			if node1.Left != nil {
				queue1 = append(queue1, node1.Left)
			}
			if node1.Right != nil {
				queue1 = append(queue1, node1.Right)
			}
			if node2.Left != nil {
				queue2 = append(queue2, node2.Left)
			}
			if node2.Right != nil {
				queue2 = append(queue2, node2.Right)
			}
		}

	}
	return true
}
