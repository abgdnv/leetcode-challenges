package main

func minSubarray(nums []int, p int) int {
	remainder := 0
	for _, v := range nums {
		if v%p != 0 {
			remainder = (remainder + v) % p
		}
	}
	if remainder == 0 {
		return 0
	}

	subArrSize := len(nums)
	sum := 0
	m := make(map[int]int)
	m[0] = -1
	for i := range len(nums) {
		sum = (sum + nums[i]) % p
		want := (sum - remainder + p) % p
		if subRemIdx, exists := m[want]; exists {
			subArrSize = min(subArrSize, i-subRemIdx)
		}
		m[sum] = i
	}
	if subArrSize != len(nums) {
		return subArrSize
	}
	return -1
}
