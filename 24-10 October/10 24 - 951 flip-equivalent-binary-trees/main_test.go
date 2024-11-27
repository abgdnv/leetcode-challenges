package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		root1 *TreeNode
		root2 *TreeNode
		want  bool
	}{
		{
			&TreeNode{
				1,
				&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{7, nil, nil}, &TreeNode{8, nil, nil}}},
				&TreeNode{3, &TreeNode{6, nil, nil}, nil},
			},
			&TreeNode{
				1,
				&TreeNode{3, &TreeNode{6, nil, nil}, nil},
				&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{7, nil, nil}, &TreeNode{8, nil, nil}}},
			},
			true,
		},
		{
			nil,
			nil,
			true,
		},
		{
			&TreeNode{1, nil, nil},
			nil,
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := flipEquiv(tC.root1, tC.root2)
			if value != tC.want {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
