package _079_word_search

// 单词搜索
// https://leetcode.cn/problems/word-search/

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if process(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

// 从(i, j)位置开始, index 表示从path[index...]是否可以在board上找到符合的路径
func process(board [][]byte, path string, index, i, j int) bool {
	if index == len(path) { // 收集答案
		return true
	}
	// 越界(剪枝)
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}
	// 不符合要求(剪枝)
	if board[i][j] != path[index] {
		return false
	}

	// 标记现场
	tmp := board[i][j]
	board[i][j] = 0
	// 进入下层决策(上、右、下、左)
	flag := process(board, path, index+1, i-1, j) ||
		process(board, path, index+1, i, j+1) ||
		process(board, path, index+1, i+1, j) ||
		process(board, path, index+1, i, j-1)

	// 恢复现场
	board[i][j] = tmp
	return flag
}
