package _054_spiral_matrix

// 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/

// 思路: 搞一个左上角点和右下角点, 两个点同时向中间靠拢, 一层层打印(矩形需要注意可能会出现是中间层为横线或竖线)
var result []int

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || matrix[0] == nil || len(matrix[0]) == 0 {
		return nil
	}
	result = make([]int, 0)
	// 左上角点(a, b); 右下角点(c, d)
	a := 0
	b := 0
	c := len(matrix) - 1
	d := len(matrix[0]) - 1
	for a <= c && b <= d {
		addEdge(matrix, a, b, c, d)
		a++
		b++
		c--
		d--
	}
	return result
}

func addEdge(matrix [][]int, a, b, c, d int) {
	if a == c { // 横线
		for i := b; i <= d; i++ {
			result = append(result, matrix[a][i])
		}
	} else if b == d { // 竖线
		for i := a; i <= c; i++ {
			result = append(result, matrix[i][b])
		}
	} else { // 矩形环(左上角开始)
		curR := a
		curC := b
		for curC != d { // top, 结束后在右上角点
			result = append(result, matrix[a][curC])
			curC++
		}
		for curR != c { // right, 结束后在右下角点
			result = append(result, matrix[curR][d])
			curR++
		}
		for curC != a { // low, 结束后在左下角点
			result = append(result, matrix[c][curC])
			curC--
		}
		for curR != a { // left, 结束后在右上角点
			result = append(result, matrix[curR][b])
			curR--
		}
	}
}
