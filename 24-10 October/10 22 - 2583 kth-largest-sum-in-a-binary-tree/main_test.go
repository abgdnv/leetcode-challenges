package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		root TreeNode
		k    int
		want int64
	}{
		{
			TreeNode{411310, &TreeNode{211244, nil, nil}, &TreeNode{111674, nil, nil}},
			2,
			322918,
		},
		{
			TreeNode{5, &TreeNode{8, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}, &TreeNode{1, nil, nil}}, &TreeNode{9, &TreeNode{3, nil, nil}, &TreeNode{7, nil, nil}}},
			2,
			13,
		},
		{
			TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, nil},
			1,
			3,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := kthLargestLevelSum(&tC.root, tC.k)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
