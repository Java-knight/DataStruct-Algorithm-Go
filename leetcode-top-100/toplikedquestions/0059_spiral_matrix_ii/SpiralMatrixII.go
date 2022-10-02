package _059_spiral_matrix_ii

// 螺旋矩阵II
// https://leetcode.cn/problems/spiral-matrix-ii/

// 思路: 搞一个左上角点和右下角点, 两个点同时向中间靠拢, 一层层打印(nxn肯定是正方形)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	a := 0
	b := 0
	c := n - 1
	d := n - 1
	start := 1
	end := n * n
	for a <= c && b <= d {
		start = addEdge(matrix, a, b, c, d, start, end)
		a++
		b++
		c--
		d--
	}
	return matrix
}

func addEdge(matrix [][]int, a, b, c, d, start, end int) int {
	if start == end { // matrix已经遍历结束了(填写最中间的)
		matrix[a][b] = start
	} else {
		curR := a
		curC := b
		for curC != d { // top, 结束后在右上角点
			matrix[a][curC] = start
			start++
			curC++
		}
		for curR != d { // right, 结束后在右下角点
			matrix[curR][d] = start
			start++
			curR++
		}
		for curC != b { // down, 结束后在左下角点
			matrix[c][curC] = start
			start++
			curC--
		}
		for curR != a { // left, 结束后在左上角点
			matrix[curR][b] = start
			start++
			curR--
		}
	}
	return start
}
