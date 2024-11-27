package main

func canArrange(arr []int, k int) bool {
	m := make(map[int]int)
	for i := range len(arr) {
		m[(arr[i]%k+k)%k]++
	}
	for i := range len(arr) {
		rem := (arr[i]%k + k) % k
		if rem == 0 {
			if m[rem]%2 != 0 {
				return false
			}
		} else {
			if m[rem] != m[k-rem] {
				return false
			}
		}
	}
	return true
}
