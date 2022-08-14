## [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/)

```go
// go 语言封装三目运算符(还是有很大不同，甚用，代码优化可以使用，确保传进来的b、c都是显式的值)
func ternary(a bool, b, c interface{}) interface{} {
    if a {
        return b
    }
    return c
}
```
