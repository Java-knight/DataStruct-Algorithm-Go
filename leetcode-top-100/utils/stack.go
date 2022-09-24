package utils

import "errors"

type Stack []interface{}

// Push 压入一个元素
func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

// Pop 弹出一个元素
func (stack *Stack) Pop() (interface{}, error) {
	val, err := stack.Peek()
	*stack = (*stack)[:len(*stack)-1]
	return val, err
}

// Peek 已经处理 error
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack empty")
	}
	return (*stack)[len(*stack)-1], nil
}

// Size 返回当前stack中的元素个数
func (stack *Stack) Size() int {
	return len(*stack)
}

// IsEmpty stack 是否为空
func (stack *Stack) IsEmpty() bool {
	if len(*stack) == 0 {
		return true
	}
	return false
}
