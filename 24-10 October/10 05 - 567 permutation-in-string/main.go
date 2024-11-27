package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	hm1 := make([]int, 26)
	hm2 := make([]int, 26)

	for _, v := range s1 {
		hm1[v-'a']++
	}

	l := 0
	r := len(s1) - 1

	for i := 0; i < len(s1); i++ {
		hm2[s2[i]-'a']++
	}
	if check(hm1, hm2) {
		return true
	}
	for r < len(s2)-1 {
		l++
		r++
		hm2[s2[r]-'a']++
		hm2[s2[l-1]-'a']--
		if check(hm1, hm2) {
			return true
		}
	}
	return false
}

func check(hm1 []int, hm2 []int) bool {
	for i := range 26 {
		if hm1[i] > hm2[i] {
			return false
		}
	}
	return true
}
