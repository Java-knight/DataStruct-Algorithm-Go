package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_List2Ints(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{}, List2Ints(nil), "输入 nil，没有返回 []int{}")

	// 1 —> 2 —> 3
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	ast.Equal([]int{1, 2, 3}, List2Ints(list), "没有成功转换成 []int")

	limit := 100
	// int [0,0,0,...]
	overLimitList := Ints2List(make([]int, limit+1))
	ast.Panics(func() {
		List2Ints(overLimitList)
	}, "转换深度超过 %d 限制的链表，没有 panic", limit)
}

func Test_Ints2List(t *testing.T) {
	ast := assert.New(t)
	ast.Nil(Ints2List([]int{}), "输入 []int{}，没有返回 nil")

	list := Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	index := 1
	for list != nil {
		ast.Equal(index, list.Val, "对应的值不对")
		list = list.Next
		index++
	}
}

func Test_GetNodeWith(t *testing.T) {
	ast := assert.New(t)

	list := Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	val := 10
	node := &ListNode{
		Val: val,
	}
	tail := list
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	expected := node                // 期望的node（正确的结果）
	actual := list.GetNodeWith(val) // 实际node（自己的方法得到的结果）
	ast.Equal(expected, actual)
}

func Test_Ints2ListWithCycle(t *testing.T) {
	ast := assert.New(t)
	arr := []int{1, 2, 3}
	noCycleList := Ints2ListWithCycle(arr, -1)
	ast.Equal(arr, List2Ints(noCycleList))

	cycleList := Ints2ListWithCycle(arr, 1)
	ast.Panics(func() {
		List2Ints(cycleList)
	})
}
