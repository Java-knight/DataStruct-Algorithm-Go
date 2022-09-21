package main

import "log"

// 题目：如何不用额外变量交换两个数
func swap() {

}

func main() {
	a := 16
	b := 603
	log.Println("交换前: ")
	log.Println(a)
	log.Println(b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	log.Println("交换后: ")
	log.Println(a)
	log.Println(b)

	arr := []int{3, 1, 100}
	i := 0
	j := 0

	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]

	log.Println(arr[i], ", ", arr[j])
}
