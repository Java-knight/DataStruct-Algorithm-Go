package _130_surrounded_reqions

// 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/

// 思路: 将board看成是一个水坝, 四边如果是 'O' 的地方就能进水, 水只能上下左右的流,
//      'X'就是石头, 遇到石头就停止, 被石头包围的水最终也会变成石头。
// 代码实现: 先将四周水能流到的位置'O' 改为 'F'(染色/保护); 然后将整个board进行还原, 将'O'改为'X', 将'F'改为'O'
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	row := len(board)
	col := len(board[0])
	for j := 0; j < col; j++ {
		if board[0][j] == 'O' { // top
			free(board, 0, j)
		}
		if board[row-1][j] == 'O' { // down
			free(board, row-1, j)
		}
	}
	for i := 0; i < row; i++ {
		if board[i][0] == 'O' { // left
			free(board, i, 0)
		}
		if board[i][col-1] == 'O' { // right
			free(board, i, col-1)
		}
	}

	// 现在将board中的 'F'改为'O'(将该位置的'O' 保护起来), 将'O'改为'X'
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}
}

func free(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != 'O' { // 越界/遇到石头
		return
	}
	board[i][j] = 'F'
	free(board, i-1, j) // 上
	free(board, i, j+1) // 右
	free(board, i+1, j) // 下
	free(board, i, j-1) // 左
}
