package _084_largest_rectangle_in_histogram

// 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

// 思路: 和42接雨水很像, 但这个需要计算每个height的最近左边和右边的最小值
// 使用单调栈(大压小)来计算process(arr, index) 计算arr中index位置左右两边离index的最小值(或等于arr[index])
// 比如: arr[4, 8, 2, 3, 5, 5, 7], index = 5 ——> leftMin=3, rightMin=5
// 注意: 第一个for中的for中等于弹出算出来可能不是最正确解, 所以后面还有个最后一个for计算最终的结果
func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	// 存放的是下标
	stack := make([]int, 0)
	result := 0
	for i, val := range heights {
		// 这里是小于等于的时候都弹出计算结果, 等于的时候可能结果不是正确的
		// !stack.isEmpty() && val <= heights[stack.peek()]
		for len(stack) != 0 && val <= heights[stack[len(stack)-1]] {
			// stack.pop()
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// k = stack.isEmpty() ? -1 : stack.peek()
			var k int
			if len(stack) == 0 {
				k = -1
			} else {
				k = stack[len(stack)-1]
			}
			// 这个计算结果如果是小于进来的，肯定正确; 等于进来的可能不正确, 但是后面会再次计算, 最终结果是正确的
			curArea := (i - k - 1) * heights[j]
			result = Max(result, curArea)
		}
		stack = append(stack, i)
	}

	// stack 中是大压小, 可以解决最终结果正确
	// 比如case: heights{5, 5, 5, 5, 5} 现在栈中放的是4下标; k=-1, curArea=(5-(-1)-1)*5=25
	// 结论: index位置左边的下标（距离index位置左边最小的位置）是正确的, 而右边的位置等遍历完arr, 也会得到正确的右边下标(距离index位置右边最小的位置)
	for len(stack) != 0 {
		// stack.pop()
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// k = stack.isEmpty() ? -1 : stack.peek()
		var k int
		if len(stack) == 0 {
			k = -1
		} else {
			k = stack[len(stack)-1]
		}
		curArea := (len(heights) - k - 1) * heights[j]
		result = Max(result, curArea)
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
