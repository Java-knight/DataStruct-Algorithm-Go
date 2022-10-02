package _051_n_queens

// N 皇后
// https://leetcode.cn/problems/n-queens/

var result [][]string

func solveNQueens(n int) [][]string {
	result = make([][]string, 0)
	path := make([][]byte, n)
	for i := range path {
		path[i] = make([]byte, n)
		for j := range path[i] {
			path[i][j] = '.'
		}
	}

	process(path, 0)
	return result
}

func process(path [][]byte, row int) {
	if row == len(path) { // 收集答案
		result = append(result, conversion(path))
	} else {
		for col := 0; col < len(path); col++ {
			if !visited(path, row, col) {
				continue
			}
			// 标记现场
			path[row][col] = 'Q'
			// 进入下层决策
			process(path, row+1)
			// 恢复现场
			path[row][col] = '.'
		}
	}
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

func conversion(path [][]byte) []string {
	ans := make([]string, len(path))
	for i := range ans {
		ans[i] = string(path[i])
	}
	return ans
}
