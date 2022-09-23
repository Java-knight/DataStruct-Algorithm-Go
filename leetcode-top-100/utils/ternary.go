package utils

// go语言对三目运算符的支持
// 注意: 传入func的值, 调用方一定要保证不能发生panic

func TernaryBool(flag bool, a, b bool) bool {
	if flag {
		return a
	}
	return b
}

func TernaryInt(flag bool, a, b int) int {
	if flag {
		return a
	}
	return b
}

func TernaryInt8(flag bool, a, b int8) int8 {
	if flag {
		return a
	}
	return b
}

func TernaryInt16(flag bool, a, b int16) int16 {
	if flag {
		return a
	}
	return b
}

func TernaryInt32(flag bool, a, b int32) int32 {
	if flag {
		return a
	}
	return b
}

func TernaryInt64(flag bool, a, b int64) int64 {
	if flag {
		return a
	}
	return b
}

// TernaryByte byte的就是 unint8
func TernaryByte(flag bool, a, b byte) byte {
	if flag {
		return a
	}
	return b
}

// TernaryRune rune 的就是 int32
func TernaryRune(flag bool, a, b rune) rune {
	if flag {
		return a
	}
	return b
}

func TernaryString(flag bool, a, b string) string {
	if flag {
		return a
	}
	return b
}
