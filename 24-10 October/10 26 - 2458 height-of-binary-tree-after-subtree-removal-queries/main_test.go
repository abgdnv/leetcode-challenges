package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		root *TreeNode
		k    []int
		want []int
	}{
		{
			&TreeNode{1,
				&TreeNode{3, &TreeNode{2, nil, nil}, nil},
				&TreeNode{4, &TreeNode{6, nil, nil}, &TreeNode{5, nil, &TreeNode{7, nil, nil}}}},
			[]int{4},
			[]int{2},
		},
		{
			&TreeNode{5,
				&TreeNode{8,
					&TreeNode{2,
						&TreeNode{4, nil, nil},
						&TreeNode{6, nil, nil}},
					&TreeNode{1, nil, nil}},
				&TreeNode{9,
					&TreeNode{3, nil, nil},
					&TreeNode{7, nil, nil}}},
			[]int{3, 2, 4, 8},
			[]int{3, 2, 3, 2},
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintln(i, tC.want), func(t *testing.T) {
			value := treeQueries(tC.root, tC.k)
			if !reflect.DeepEqual(value, tC.want) {
				t.Fatalf(`got %v, expected %v`, value, tC.want)
			}
		})
	}
}
