package _037_sudoku_solver

// 解数独
// https://leetcode.cn/problems/sudoku-solver/

func solveSudoku(board [][]byte) {
	row := make([][]bool, 9)
	col := make([][]bool, 9)
	bucket := make([][]bool, 9)
	for i := range row {
		row[i] = make([]bool, 10)
	}
	for i := range col {
		col[i] = make([]bool, 10)
	}
	for i := range bucket {
		bucket[i] = make([]bool, 10)
	}
	initMaps(board, row, col, bucket)
	process(board, 0, 0, row, col, bucket)
}

// 初始化row、col、bucket, 将目前board的状态设置进去
func initMaps(board [][]byte, row, col, bucket [][]bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			bid := 3*(i/3) + (j / 3) // 3x3的小方格位置
			if board[i][j] != '.' {  // 不为数字时设置值
				num := board[i][j] - '0'
				row[i][num] = true
				col[j][num] = true
				bucket[bid][num] = true
			}
		}
	}
}

// DFS + 剪支
func process(board [][]byte, i, j int, row, col, bucket [][]bool) bool {
	if i == 9 { // 收集答案
		return true
	}
	nexti := TernaryInt(j != 8, i, i+1)
	nextj := TernaryInt(j != 8, j+1, 0)
	if board[i][j] != '.' { // 已经是num, 没有操作性进入下层决策(剪支1)
		return process(board, nexti, nextj, row, col, bucket)
	} else { // 是'.', 可以进行操作
		bid := 3*(i/3) + (j / 3)
		for num := 1; num <= 9; num++ {
			// 只有三个判断条件都是false的情况下进入(剪支2)
			if (!row[i][num]) && (!col[j][num]) && (!bucket[bid][num]) {
				// 标记现场
				board[i][j] = byte(num - '0')
				row[i][num] = true
				col[j][num] = true
				bucket[bid][num] = true
				// 进入下层决策
				if process(board, nexti, nextj, row, col, bucket) {
					return true
				}
				// 恢复现场
				board[i][j] = '.'
				row[i][num] = false
				col[j][num] = false
				bucket[bid][num] = false
			}
		}
		return false
	}
}

func TernaryInt(flag bool, i, j int) int {
	if flag {
		return i
	}
	return j
}
