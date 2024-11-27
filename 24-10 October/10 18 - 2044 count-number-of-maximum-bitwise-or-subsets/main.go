package main

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	maxOrValue := 0

	for _, num := range nums {
		maxOrValue |= num
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, maxOrValue+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return countSubsetsRecursive(nums, 0, 0, maxOrValue, memo)
}

func countSubsetsRecursive(nums []int, index, currentOr, targetOr int, memo [][]int) int {
	// Base case: reached the end of the array
	if index == len(nums) {
		if currentOr == targetOr {
			return 1
		}
		return 0
	}

	// Check if the result for this state is already memoized
	if memo[index][currentOr] != -1 {
		return memo[index][currentOr]
	}

	// Don't include the current number
	countWithout := countSubsetsRecursive(nums, index+1, currentOr, targetOr, memo)

	// Include the current number
	countWith := countSubsetsRecursive(nums, index+1, currentOr|nums[index], targetOr, memo)

	// Memoize and return the result
	memo[index][currentOr] = countWithout + countWith
	return memo[index][currentOr]
}
