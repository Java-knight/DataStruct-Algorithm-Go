package _045_jump_game_ii

// 跳跃游戏II
// https://leetcode.cn/problems/jump-game-ii/

// 思想: 上帝视角, 提前就可以直到你 步数+1 的最远距离
func jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	result := 0     // 最小跳跃步数
	cur := 0        // 当前result步可以到达的最远距离
	next := nums[0] // result+1可以达到的最远距离
	for i := 1; i < len(nums); i++ {
		if i > cur {
			result++
			cur = next
		}
		next = Max(next, i+nums[i])
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

//func jump(nums []int) int {
//	if nums == nil || len(nums) == 0 {
//		return 0
//	}
//	result := 0  // 最小跳跃步数
//	cur := 0  // 当前result步可以到达的最远距离
//	next := nums[0]  // result+1可以达到的最远距离
//	for i := 1; i < len(nums); i++ {
//		if next >= len(nums) - 1 {  // next已经超过了数组最大了, 直接return
//			return result + 1
//		}
//		if i > cur {
//			result++
//			cur = next
//		}
//		next = Max(next, i + nums[i])
//	}
//	return result
//}
