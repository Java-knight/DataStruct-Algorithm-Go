package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 问题集合
type question1 struct {
	name   string // 名字
	param1        // 参数
	ans1          // 期望值
}

// 参数
type param1 struct {
	nums   []int
	target int
}

// 正确结果
type ans1 struct {
	realResult []int
}

func Test_Problem1(t *testing.T) {
	tests := []question1{
		{
			name:   "base case1",
			param1: param1{[]int{3, 2, 4}, 6},
			ans1:   ans1{[]int{1, 2}},
		},
		{
			name:   "base case2",
			param1: param1{[]int{3, 2, 4}, 5},
			ans1:   ans1{[]int{0, 1}},
		},
		{
			name:   "base case3",
			param1: param1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1:   ans1{[]int{1, 3}},
		},
		{
			name:   "base case4",
			param1: param1{[]int{0, 1}, 1},
			ans1:   ans1{[]int{0, 1}},
		},
		{
			name:   "base case5",
			param1: param1{[]int{0, 3}, 5},
			ans1:   ans1{nil},
		},
		// 也可以增加 cast
	}

	fmt.Printf("--------------------------Leetcode Prodlem 1--------------------------\n")
	fmt.Printf("--------------------------------两数之和--------------------------------\n")
	fmt.Printf("------------------https://leetcode.cn/problems/two-sum/ -----------------\n")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.param1.nums, tt.param1.target)
			assert.Equal(t, tt.ans1.realResult, result)
			fmt.Printf("[input]: %v       [output]: %v\n", tt.param1, result)
		})
	}
	fmt.Printf("\n\n\n")
}
