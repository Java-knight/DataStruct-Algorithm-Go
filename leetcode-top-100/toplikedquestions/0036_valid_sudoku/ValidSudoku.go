package _036_valid_sudoku

// 有效的数独
// https://leetcode.cn/problems/valid-sudoku/

// 暴力判断
func isValidSudoku(board [][]byte) bool {
	// row[i][num] 表示在第i行num是否存在, col同理
	// 因为num是0~9, row[i][0]位置不用, 方便运算, 所以是10个
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

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			bid := 3*(i/3) + (j / 3) // 小3x3宫格位置
			if board[i][j] != '.' {  // 是数字进入判断
				num := board[i][j] - '0' // 这里就实现了row\col设置10的方便了
				// 3个判断条件
				if row[i][num] || col[j][num] || bucket[bid][num] {
					return false
				}
				// 设置值
				row[i][num] = true
				col[j][num] = true
				bucket[bid][num] = true
			}
		}
	}
	return true
}
