package main

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
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
func main() {

}
