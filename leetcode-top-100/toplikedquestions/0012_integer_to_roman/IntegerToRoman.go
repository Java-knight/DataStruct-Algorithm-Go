package _012_integer_to_roman

import (
	"bytes"
)

// 整数转罗马数字
// https://leetcode.cn/problems/integer-to-roman/

func intToRoman(num int) string {
	temp := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}

	sb := NewStringBuilderNoParam()
	sb.Append(temp[3][num/1000%10]).Append(temp[2][num/100%10]).Append(temp[1][num/10%10]).Append(temp[0][num%10])
	return sb.String()
}

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilderNoParam() *StringBuilder {
	return &StringBuilder{}
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) String() string {
	return builder.buffer.String()
}

// 可以使用utils 方法
//func intToRoman(num int) string {
//	temp := [][]string{
//		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
//		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
//		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
//		{"", "M", "MM", "MMM"},
//	}
//
//	sb := utils.NewStringBuilderNoParam()
//	sb.Append(temp[3][num / 1000 % 10]).Append(temp[2][num / 100 % 10]).Append(temp[1][num / 10 % 10]).Append(temp[0][num % 10])
//	return sb.String()
//}
