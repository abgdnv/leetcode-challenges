package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	nodeDepths := make(map[int]int)
	subtreeHeights := make(map[int]int)

	firstLargestHeight := make(map[int]int)
	secondLargestHeight := make(map[int]int)

	dfs(
		root,
		0,
		nodeDepths,
		subtreeHeights,
		firstLargestHeight,
		secondLargestHeight,
	)

	results := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		queryNode := queries[i]
		nodeLevel := nodeDepths[queryNode]

		if subtreeHeights[queryNode] == firstLargestHeight[nodeLevel] {
			results[i] = nodeLevel + secondLargestHeight[nodeLevel] - 1
		} else {
			results[i] = nodeLevel + firstLargestHeight[nodeLevel] - 1
		}
	}
	return results
}

func dfs(
	node *TreeNode,
	level int,
	nodeDepths,
	subtreeHeights,
	firstLargestHeight,
	secondLargestHeight map[int]int,
) int {

	if node == nil {
		return 0
	}
	nodeDepths[node.Val] = level

	// Calculate the height of the current subtree
	leftHeight := dfs(
		node.Left,
		level+1,
		nodeDepths,
		subtreeHeights,
		firstLargestHeight,
		secondLargestHeight)
	rightHeight := dfs(
		node.Right,
		level+1,
		nodeDepths,
		subtreeHeights,
		firstLargestHeight,
		secondLargestHeight)
	currentHeight := 1 + max(leftHeight, rightHeight)

	subtreeHeights[node.Val] = currentHeight

	// Update the largest and second largest heights at the current level
	currentFirstLargest := firstLargestHeight[level]
	if currentHeight > currentFirstLargest {
		secondLargestHeight[level] = currentFirstLargest
		firstLargestHeight[level] = currentHeight
	} else if currentHeight > secondLargestHeight[level] {
		secondLargestHeight[level] = currentHeight
	}

	return currentHeight
}
