package _052_n_queens_ii

// N 皇后II
// https://leetcode.cn/problems/n-queens-ii/

func totalNQueens(n int) int {
	path := make([][]byte, n)
	for i := range path {
		path[i] = make([]byte, n)
		for j := range path[i] {
			path[i][j] = '.'
		}
	}
	return process(0, path, 0)
}

func process(result int, path [][]byte, row int) int {
	if row == len(path) { // 收集答案
		result++
	} else {
		for col := 0; col < len(path); col++ {
			if !visited(path, row, col) {
				continue
			}
			// 标记现场
			path[row][col] = 'Q'
			// 进入下层决策
			result = process(result, path, row+1)
			// 恢复现场
			path[row][col] = '.'
		}
	}
	return result
}

func visited(path [][]byte, row, col int) bool {
	size := len(path)
	var i, j int
	// 正上方
	for i = 0; i < row; i++ {
		if path[i][col] == 'Q' {
			return false
		}
	}

	// 左上角
	i = row - 1
	j = col - 1
	for i >= 0 && j >= 0 {
		if path[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	// 右上角
	i = row - 1
	j = col + 1
	for i >= 0 && j < size {
		if path[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
