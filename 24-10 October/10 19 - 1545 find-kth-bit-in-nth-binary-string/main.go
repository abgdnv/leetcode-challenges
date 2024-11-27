package main

func findKthBit(n int, k int) byte {
	invertCount := 0
	len := (1 << n) - 1 // Length of Sn is 2^n - 1

	for k > 1 {
		// If k is in the middle, return based on inversion count
		if k == len/2+1 {
			if invertCount%2 == 0 {
				return '1'
			} else {
				return '0'
			}
		}

		// If k is in the second half, invert and mirror
		if k > len/2 {
			k = len + 1 - k // Mirror position
			invertCount++   // Increment inversion count
		}

		len /= 2 // Reduce length for next iteration
	}

	// For the first position, return based on inversion count
	if invertCount%2 == 0 {
		return '0'
	} else {
		return '1'
	}
}

//simulation
// func findKthBit(n int, k int) byte {
// 	s := make([]string, n+1)
// 	s[0] = "0"
// 	for i := 1; i <= n; i++ {
// 		s[i] = s[i-1] + "1" + Reverse(Invert(s[i-1]))
// 	}
// 	return s[n][k-1]
// }

// func Reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func Invert(s string) string {
// 	l := len(s)
// 	runes := make([]rune, l)
// 	for i := range s {
// 		if rune(s[i]) == '0' {
// 			runes[i] = '1'
// 		} else {
// 			runes[i] = '0'
// 		}
// 	}
// 	return string(runes)
// }
