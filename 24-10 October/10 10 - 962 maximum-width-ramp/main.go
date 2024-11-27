package main

func main() {
	maxWidthRamp([]int{6, 0, 8, 2, 1, 5}) //4
}

func maxWidthRamp(nums []int) int {
	n := len(nums)
	rightMax := make([]int, n)

	// Fill rightMax array with the maximum values from the right
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	left, right := 0, 0
	maxWidth := 0

	// Traverse the array using left and right pointers
	for right < n {
		// Move left pointer forward if current left exceeds rightMax
		for left < right && nums[left] > rightMax[right] {
			left++
		}
		maxWidth = max(maxWidth, right-left)
		right++
	}

	return maxWidth
}

// O(nlog(n))
// func maxWidthRamp(nums []int) int {
// 	idx := make([]int, len(nums))
// 	for i := range idx {
// 		idx[i] = i
// 	}
// 	slices.SortFunc(idx, func(i int, j int) int {
// 		if nums[i] != nums[j] {
// 			return nums[i] - nums[j]
// 		}
// 		return i - j
// 	})
// 	minIdx := len(nums)
// 	maxl := 0
// 	for _, i := range idx {
// 		maxl = max(maxl, i-minIdx)
// 		minIdx = min(minIdx, i)
// 	}
// 	return maxl
// }

//naive
// func maxWidthRamp(nums []int) int {
// 	maxl := 0
// 	for i := range len(nums) {
// 		l := 0
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] <= nums[j] {
// 				l = max(l, j-i)
// 			}
// 		}
// 		maxl = max(l, maxl)
// 	}
// 	return maxl
// }
