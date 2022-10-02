package _118_pascals_triangle

// 杨辉三角
// https://leetcode.cn/problems/pascals-triangle/

/**
 * 杨辉三角, 当 level 大于2时, 左右两边都是1, 而中间的值是由自己上方+左上方的和
 *  nums[i][j] = nums[i - 1][j] + nums[i - 1][j - 1]
 * 1
 * 1  1
 * 1  2  1
 * 1  3  3  1
 * 1  4  6  4  1
 */
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	// (1) nums[i][0] = 1
	for i := range result {
		result[i] = make([]int, i+1)
		result[i][0] = 1
	}
	// (2) 将每level 的中间值和左右边的值填充
	for i := 1; i < numRows; i++ {
		for j := 1; j < i; j++ { // nums[i][j] = nums[i-1][j] + nums[i-1][j-1]
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
		result[i][i] = 1 // nums[i][i] = 1
	}
	return result
}
