package utils

// 布尔类型的异或操作
func XORBool(a, b bool) bool {
	return (a || b) && !(a && b)
}
