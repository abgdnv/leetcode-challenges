package main

func countSquares(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row+1)
	for i := range row + 1 {
		dp[i] = make([]int, col+1)
	}

	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i+1][j+1] = min(
					min(dp[i][j+1], dp[i+1][j]),
					dp[i][j],
				) + 1
				ans += dp[i+1][j+1]
			}
		}
	}
	return ans
}

// func countSquares(matrix [][]int) int {
// 	count := 0
// 	lengthR := len(matrix)
// 	lengthC := len(matrix[0])
// 	for i := range lengthR {
// 		for j := range lengthC {
// 			if matrix[i][j] == 1 {
// 				count++
// 				for l := 1; j+l < lengthC && i+l < lengthR; l++ {
// 					if check(matrix, i, j, l) {
// 						count++
// 					} else {
// 						break
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// func check(matrix [][]int, x, y, l int) bool {
// 	for i := x; i <= x+l; i++ {
// 		for j := y; j <= y+l; j++ {
// 			if matrix[i][j] == 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
