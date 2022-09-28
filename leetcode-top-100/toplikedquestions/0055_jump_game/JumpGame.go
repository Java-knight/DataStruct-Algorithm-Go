package _055_jump_game

// 跳跃游戏
// https://leetcode.cn/problems/jump-game/

func canJump(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return true
	}
	max := nums[0] // 表示能到达的最大下标
	for i := 1; i < len(nums); i++ {
		if i > max {
			return false
		}
		max = Max(max, i+nums[i])
	}
	return true
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//func canJump(nums []int) bool {
//	if nums == nil || len(nums) < 2 {
//		return true
//	}
//	max := nums[0]  // 表示能到达的最大下标
//	for i := 1; i < len(nums); i++ {
//		if max > len(nums) - 1 {  // 优化:已经提前直到可以达到size-1, 可以直接退出
//			break
//		}
//		if i > max {
//			return false
//		}
//		max = Max(max, i + nums[i])
//	}
//	return true
//}
