package _013_roman_to_integer

// 罗马数字转整数
// https://leetcode.cn/problems/roman-to-integer/

func romanToInt(s string) int {
	temp := make([]int, len(s))
	for i := range s {
		switch s[i] {
		case 'M':
			temp[i] = 1000
			break
		case 'D':
			temp[i] = 500
			break
		case 'C':
			temp[i] = 100
			break
		case 'L':
			temp[i] = 50
			break
		case 'X':
			temp[i] = 10
			break
		case 'V':
			temp[i] = 5
			break
		case 'I':
			temp[i] = 1
			break
		}
	}
	sum := 0
	for i := 0; i < len(temp)-1; i++ {
		if temp[i] < temp[i+1] {
			sum -= temp[i]
		} else {
			sum += temp[i]
		}
	}
	return sum + temp[len(temp)-1]
}
