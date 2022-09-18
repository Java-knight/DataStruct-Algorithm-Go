package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

// 插入排序
// 验证 https://www.nowcoder.com/practice/2baf799ea0594abd974d37139de27896?tpId=308&tqId=1089529&ru=/exam/oj&qru=/ta/algorithm-start/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D308

// 时间复杂度 O(N^2), 空间复杂度 O(1)
func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 0 ~ 1   做到有序
	// 0 ~ 2   做到有序
	// 0 ~ 3   做到有序
	// 0 ~ N-1 做到有序
	for i := 1; i < len(arr); i++ { // 0 ~ i 保证有序
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 对数器
func comparator(arr []int) {
	// sort.Slice() 不稳定的排序，底层使用快排
	// func SliceStable() 稳定的排序，底层使用快排
	sort.Slice(arr[:], func(i, j int) bool {
		return arr[i] < arr[j]
	})
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

// for test
func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	size := len(arr)
	result := make([]int, size)
	for i, val := range arr {
		result[i] = val
	}
	return result
}

// for test
func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		insertionSort(arr1)
		comparator(arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
			log.Println(arr1)
			log.Println(arr2)
			break
		}
	}

	if succeed {
		log.Println("Nick!")
	} else {
		log.Println("Fucking fucked!")
	}

	arr := generateRandomArray(maxSize, maxValue)
	log.Println(arr)
	insertionSort(arr)
	log.Println(arr)
}
