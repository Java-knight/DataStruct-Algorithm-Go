package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

// 题目：在一个有序数组中, 找到 <= 某个数最右侧的位置

func nearestIndex(sortedArr []int, target int) int {
	left := 0
	right := len(sortedArr) - 1
	mid := 0
	index := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if sortedArr[mid] >= target {
			index = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return index
}

// 对数器
func comparator(sortedArr []int, target int) int {
	for i := 0; i < len(sortedArr); i++ {
		if sortedArr[i] >= target {
			return i
		}
	}
	return -1
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
		if nearestIndex(arr, value) != comparator(arr, value) {
			succeed = false
			log.Println("arr: ", arr, "target: ", value)
			log.Println("left nearestIndex result: ", nearestIndex(arr, value))
			log.Println("comparator result: ", comparator(arr, value))
			break
		}
	}

	if succeed {
		log.Println("Nick!")
	} else {
		log.Println("Fucking fucked!")
	}
}
