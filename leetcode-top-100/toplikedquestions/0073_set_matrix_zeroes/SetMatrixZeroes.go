package _073_set_matrix_zeroes

// 矩阵置零
// https://leetcode.cn/problems/set-matrix-zeroes/

// 一个布尔变量标识(面试吹牛用, 极客精神)
// (1) 这个时候指用一个 colFlag表示第0列在改变之前是否有0, 但是(0, 0)位置的数不决定第0列是否有0;
// 第0行是否有0用(0, 0)位置表示, 如果有就将matrix[0][0]改成0, 没有保持不变就行;
// (2) 接下来收集matrix[i][j]位置的0(i和j都是从1开始), 将其填入第0行和第0列
// (3) 判断第0行和第0列中是0的行和列, 将matrix出现行和列的位置填充0
// (4) 判断matrix(0, 0)位置决定第0行是否全部为0, 而colFlag决定第0列是否全部为0
// 注意: (3)中是从matrix.length-1的位置开始判断, 因为如果从0位置开始, 可能第0行在未改变前有一个0,
//       所以将(0, 0)位置改为0, 先判断0位置会导致, 将第0行和第0列全部置0的
func setZeroes(matrix [][]int) {
	colFlag := false
	i := 0
	j := 0
	for i = 0; i < len(matrix); i++ { // (1)和(2)一起进行, 只是(0, 0)位置表示第0行, 而第0列使用colFlag表示
		for j = 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j == 0 { // (1) 判断第0列是否出现0
					colFlag = true
				} else {
					matrix[0][j] = 0
				}
			}
		}
	}

	// (4) 必须先行, 后列; 判断第0行改变前是否需要全0
	for i = len(matrix) - 1; i >= 0; i-- { // (3) 将第0行和第0列的填充到matrix中(第0行是从右向左)
		for j = 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if colFlag { // (4) 处理第0列(需要通过标记为来判断原本是否需要全0)
		for i = 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

//// 方法1:
//// (1) 使用两个变量保存matrix的第0行和第0列的状态, 如果有一个是0, 标记位设置为true;
//// (2) 将matrix[i][j] == 0(i和j都是从1开始), 就记录到matrix[0][j]和matrix[i][0]位置都改成0;
//// (3) 查看matrix第0行中是0的位置, 该列全部改为0, matrix第0列中是0的位置, 该列全部改为0;
//// (4) 判断(1)中的标记位, 是否需要将matrix的第0行和第0列全部改为0
//func setZeroes(matrix [][]int)  {
//	rowFlag := false
//	colFlag := false
//	i := 0
//	j := 0
//	for i = 0; i < len(matrix); i++ {  // (1)和(2)同时进行, 将matrix[i][j] == 0的行和列记录在0行0列的位置
//		for j = 0; j < len(matrix[0]); j++ {
//			if matrix[i][j] == 0 {
//				if i == 0{  // (1) 判断第0行是否出现0
//					rowFlag = true
//				} else {
//					matrix[i][0] = 0
//				}
//				if j == 0 {  // (1) 判断第0列是否出现0
//					colFlag = true
//				} else {
//					matrix[0][j] = 0
//				}
//			}
//		}
//	}
//
//	for i = 1; i < len(matrix); i++ {  // (3) 第0行和第0列开始按行填充matrix
//		for j = 1; j < len(matrix[0]); j++ {
//			if (matrix[i][0] == 0 || matrix[0][j] == 0) {
//				matrix[i][j] = 0
//			}
//		}
//	}
//
//	if colFlag {  // (4) 处理第0列(需要通过标记为来判断原本是否需要全0)
//		for i = 0; i < len(matrix); i++ {
//			matrix[i][0] = 0
// 		}
//	}
//	if rowFlag {  // (4) 处理第0行(需要通过标记为来判断原本是否需要全0)
//		for j = 0; j < len(matrix[0]); j++ {
//			matrix[0][j] = 0
//		}
//	}
//}
