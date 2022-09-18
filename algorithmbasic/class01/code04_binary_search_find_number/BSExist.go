package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

// 二分查找（寻找一个有序数组上是否存在一个target的数）
// 测试：https://leetcode.cn/problems/binary-search/

// 时间复杂度 O(logN)，空间复杂度 O(1)
func exist(sortedArr []int, target int) bool {
	if sortedArr == nil || len(sortedArr) == 0 {
		return false
	}

	left := 0
	right := len(sortedArr) - 1
	mid := 0
	for left < right {
		mid = left + ((right - left) >> 1)
		if sortedArr[mid] == target {
			return true
		} else if sortedArr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return sortedArr[left] == target
}

// 对数器
func comparator(sortedArr []int, target int) bool {
	for _, cur := range sortedArr {
		if cur == target {
			return true
		}
	}
	return false
}

// for test
func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize) + 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// [-maxValue+2, maxValue]
		arr[i] = (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue))
	}
	return arr
}

func main() {
	testTime := 500000
	maxSize := 10
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		value := (rand.Intn(maxValue) + 1) - (rand.Intn(maxValue))
		if exist(arr, value) != comparator(arr, value) {
			succeed = false
			log.Println("arr: ", arr, "target: ", value)
			break
		}
	}

	if succeed {
		log.Println("Nick!")
	} else {
		log.Println("Fucking fucked!")
	}
}
