package leetcode

// 思路1：滑动窗口（最优解）
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 字母表（里面记录到是字母到数量）【如果是0表示，没有记录过该字母】
	var freq [127]int
	left, right, result := 0, 0, 0 // right表示下次需要判断到index

	for left < len(s) { // 从左到右
		if right < len(s) && freq[s[right]] == 0 { // 右滑
			freq[s[right]]++
			right++
		} else { // 左滑
			freq[s[left]]--
			left++
		}
		result = max(result, right-left) // 因为 right 的含义不同，所以不需要r-l+1
	}
	return result
}

// 思路2：滑动窗口-哈希桶（消耗空间/时间 更多）
// 和思路1的不同，上面是 right指针始终在右边跑；而本思路是left 始终在右边跑
func lengthOfLongestSubstring2(s string) int {
	left, right, result := 0, 0, 0
	cache := make(map[byte]int, len(s)) // map<char, stringIndex>
	for left < len(s) {
		// 重点可以通过画图理解一下 idx >= right 步骤
		if idx, ok := cache[s[left]]; ok && idx >= right { // 右滑
			right = idx + 1
		}
		cache[s[left]] = left
		left++
		result = max(result, left-right)
	}
	return result
}

// 思路3: 位图
// right指针始终在右边跑
func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 位图集合，Ascii码表是 256 个字母
	var bitSet [256]bool
	left, right, result := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] { // 位图中存在: (1)删除记录、(2)左滑
			bitSet[s[left]] = false
			left++
		} else { // 位图中不存在: (1)记录、(2)右滑
			bitSet[s[right]] = true
			right++
		}
		result = max(result, right-left)
		// 边界条件
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
