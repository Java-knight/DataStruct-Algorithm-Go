package _048_rotate_image

// 旋转图像
// https://leetcode.cn/problems/rotate-image/

// 正方形图像, 画出一个4x4就能找下标变化规律
func rotate(matrix [][]int) {
	// 正方形 matrix.length == matrix[0].length
	tLR := 0                  // 左上角(横坐标)topLeftRow
	tLC := 0                  // 左上角(纵坐标)topLeftCol
	lRR := len(matrix) - 1    // 右下角(横坐标)lowRightRow
	lRC := len(matrix[0]) - 1 // 右下角(纵坐标)lowRightCol
	for tLR < lRR {           // 解决一层一层环进行旋转
		rotateEdge(matrix, tLR, tLC, lRR, lRC) // 始终是每个正方形的左上角和右下角
		tLR++
		tLC++
		lRR--
		lRC--
	}
}

// 每层环的旋转
func rotateEdge(m [][]int, tLR, tLC, lRR, lRC int) {
	times := lRC - tLC
	tmp := 0
	for i := 0; i < times; i++ {
		tmp = m[tLR][tLC+i]           // 左上角
		m[tLR][tLC+i] = m[lRR-i][tLC] // 左下角
		m[lRR-i][tLC] = m[lRR][lRC-i] // 右下角
		m[lRR][lRC-i] = m[tLR+i][lRC] // 右上角
		m[tLR+i][lRC] = tmp
	}
}
