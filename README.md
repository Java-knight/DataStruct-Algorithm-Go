# Go数据结构与算法刷题
algorithmbasic class02 code03_km有问题

Go语言的工具包
```go
// Int32的最大值/最小值, 类似Java的Integer.MAX_VALUE/Integer.MIN_VALUE
math.MaxInt32 = 2^31 - 1
math.MinInt32  = -2^31

// 排序方法, 类似Java的Arrays.sort(arr)
sort.Ints(arr)

// 三目运算符, 请查看utils包的ternary

// StringBuilder, 请查看 utils包的stringBuilder
// Stack, 请查看 utils包的stack
// heap的使用, 请查看tests/heap.go
// Go语言中 bool值不支持异或操作, 请查看 utils包的orx
// Go语言中不支持无符号位左右移动(<<<、>>>), 但是可以通过设置变量为无符号位
// 数字转字符串，Java一般是 String.valueOf(num), Go语言需要string(num + '0')
// 拷贝数组可以使用copy(dst, src []Type)
```

leetcode 22 如果只要求数量, 那就是科特兰数
leetcode 23 需要使用小根堆, 需要自己实现一下

## Go语言的集合
go语言也是有集合的实现的, 但是它只定义了通用的规则, 具体的调用方式完全由调用方决定。
> 集合包: sdk/container/

里面有list、heap、ring的实现, 比如`heap`, 它里面就只定义了 heapify、up、down 操作。 
而调用方需要调用的Push\Pop都是需要自己实现, heap.go 提供了一些接口, 包含比较器Less。
这样更加方便, 连IsEmpty都没有需要自己实现。离谱...