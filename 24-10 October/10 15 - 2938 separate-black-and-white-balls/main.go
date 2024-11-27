package main

func minimumSteps(s string) int64 {
	var count int64
	var moves int64
	for i := range len(s) {
		if s[i] == '1' {
			count++
		} else {
			moves += count
		}
	}
	return moves
}

// func minimumSteps(s string) int64 {
// 	n := len(s)
// 	count := 0
// 	for i := range n {
// 		if s[i] == '0' {
// 			count++
// 		}
// 	}
// 	var w0, w1 int64
// 	for i := range n {
// 		if s[i] == '1' && i < count {
// 			w1 += int64(count - i)
// 		}
// 		if s[i] == '0' && i > count {
// 			w0 += int64(i - count)
// 		}
// 	}
// 	return w1 + w0
// }
