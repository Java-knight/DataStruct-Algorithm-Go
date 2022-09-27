//package _038_count_and_say
package main

import (
	"bytes"
	"log"
	"strings"
)

func main() {
	log.Println(countAndSay(4))
}

// 外观数列
// https://leetcode.cn/problems/count-and-say/

func countAndSay(n int) string {
	if n < 1 { // 边界条件
		return ""
	}
	if n == 1 { // 递归结束的条件
		return "1"
	}
	last := countAndSay(n - 1)
	sb := NewStringBuilder()
	count := 1 // 记录相同字母个数
	for i := 1; i < len(last); i++ {
		if last[i-1] == last[i] { // i-1位置上的数字和i位置上的相同, count++
			count++
		} else {
			s := string(count + '0')
			sb.Append(s)
			sb.Append(string(last[i-1]))
			count = 1
		}
	}
	sb.Append(string(count + '0'))
	sb.Append(string(last[len(last)-1:])) // 字符串获取子串[start:last], 不填就是[0, len]
	return sb.String()
}

type StringBuilder struct {
	buffer bytes.Buffer
}

func NewStringBuilder(s ...string) *StringBuilder {
	builder := StringBuilder{}
	builder.buffer.WriteString(strings.Join(s, ""))
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) String() string {
	return builder.buffer.String()
}
