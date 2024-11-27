package main

func maxMoves(grid [][]int) int {
	mgrid := make([][]int, len(grid))
	for i := range len(grid) {
		mgrid[i] = make([]int, len(grid[0]))
		for j := range len(mgrid[i]) {
			mgrid[i][j] = -1
		}
	}
	moves := 0
	for i := range len(grid) {
		moves = max(moves, countMoves(grid, i, 0, mgrid))
	}
	return moves
}

func countMoves(grid [][]int, row, col int, mgrid [][]int) int {
	if mgrid[row][col] >= 0 {
		return mgrid[row][col]
	}
	m := len(grid)
	n := len(grid[0])
	var moves [3]int
	if row > 0 && row-1 < m && col+1 < n && grid[row][col] < grid[row-1][col+1] {
		moves[0] = countMoves(grid, row-1, col+1, mgrid) + 1
	}
	if col+1 < n && grid[row][col] < grid[row][col+1] {
		moves[1] = countMoves(grid, row, col+1, mgrid) + 1
	}
	if row+1 < m && col+1 < n && grid[row][col] < grid[row+1][col+1] {
		moves[2] = countMoves(grid, row+1, col+1, mgrid) + 1
	}
	mgrid[row][col] = max(moves[0], moves[1], moves[2])
	return mgrid[row][col]
}
