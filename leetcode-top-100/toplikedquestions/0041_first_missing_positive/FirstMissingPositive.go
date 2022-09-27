package _041_first_missing_positive

// 缺失的第一个正数
// https://leetcode.cn/problems/first-missing-positive/

// 利用双指针
// 思想: left去维护一个有序的一个数组, 保证[0, left-1]上i位置上都方的i+1; right就是一个垃圾堆, 接收left不需要的数
// 4个条件: (1) nums[left] > right 仍垃圾堆; (2) nums[left] <= left 仍垃圾堆;
// (3) nums[left] == left+1,  left++（预期不变）
// (4) left+1 < nums[left] < R 那就去nums[nums[left]-1]看看上面是不是nums[left],
//     如果是就将nums[left]仍垃圾堆; 不是就swap(left, nums[left]-1), 重新看left上的数
// 仍垃圾堆都会预期变差, right就会向左移动; 其它预期不变, right不移动
func firstMissingPositive(nums []int) int {
	left := 0
	right := len(nums)
	for left < right {
		// base case3
		if nums[left] == left+1 {
			left++
		} else if nums[left] > right || nums[left] <= left || nums[nums[left]-1] == nums[left] { // base case2/1 4的(1)
			right--
			nums[left] = nums[right] // 仍垃圾堆(把垃圾堆原本好的数放在left位置上, 此时right位置就属于垃圾堆了)
		} else { // base case4的(2)
			swap(nums, left, nums[left]-1)
		}
	}
	return left + 1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
