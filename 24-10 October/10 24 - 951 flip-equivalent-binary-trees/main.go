package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	nodePairStack := make([][]*TreeNode, 1, 100)
	nodePairStack[0] = make([]*TreeNode, 2)
	nodePairStack[0][0] = root1
	nodePairStack[0][1] = root2

	for len(nodePairStack) > 0 {
		current := nodePairStack[len(nodePairStack)-1]
		nodePairStack = nodePairStack[:len(nodePairStack)-1]
		node1 := current[0]
		node2 := current[1]

		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}

		if isEqual(node1.Left, node2.Left) && isEqual(node1.Right, node2.Right) {
			nodePairStack = append(nodePairStack, []*TreeNode{node1.Left, node2.Left})
			nodePairStack = append(nodePairStack, []*TreeNode{node1.Right, node2.Right})
		} else if isEqual(node1.Left, node2.Right) && isEqual(node1.Right, node2.Left) {
			nodePairStack = append(nodePairStack, []*TreeNode{node1.Left, node2.Right})
			nodePairStack = append(nodePairStack, []*TreeNode{node1.Right, node2.Left})
		} else {
			return false
		}
	}
	return true
}

func isEqual(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil && node1.Val == node2.Val {
		return true
	}
	return false
}
