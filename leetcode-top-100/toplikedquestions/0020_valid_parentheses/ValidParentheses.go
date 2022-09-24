package _020_valid_parentheses

import "errors"

// 有效括号
// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}
	stack := &Stack{}
	for _, val := range s {
		if val == '(' || val == '{' || val == '[' {
			stack.Push(val)
		} else {
			if stack.IsEmpty() {
				return false
			}
			last, _ := stack.Pop() // 这里的error 在 leetcode 是不会出现的
			if (val == ')' && last != '(') || (val == ']' && last != '[') || (val == '}' && last != '{') {
				return false
			}
		}
	}
	return stack.IsEmpty()

}

type Stack []interface{}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

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

func (stack *Stack) Size() int {
	return len(*stack)
}

// stack 是否为空
func (stack *Stack) IsEmpty() bool {
	if len(*stack) == 0 {
		return true
	}
	return false
}
