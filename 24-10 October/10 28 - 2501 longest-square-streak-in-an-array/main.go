package main

func longestSquareStreak(nums []int) int {

	m := make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}
	maxStreake := 0

	for _, num := range nums {
		sl := 0
		n := num
		for m[n] {
			sl++
			n *= n
			if n > 1e5 {
				break
			}
		}
		maxStreake = max(maxStreake, sl)
	}

	if maxStreake < 2 {
		return -1
	}
	return maxStreake
}
