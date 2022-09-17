package code01_selection_sort

// 选择排序
// 测试：https://www.nowcoder.com/practice/2baf799ea0594abd974d37139de27896?tpId=308&tqId=1089529&ru=/exam/oj&qru=/ta/algorithm-start/question-ranking&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D308
// 时间复杂度 O(N^2), 空间复杂度 O(1)
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// 0 ~ N-1 找到最小值，放到 0 位置上
	// 1 ~ N-1 找到最小值，放到 1 位置上
	// 2 ~ N-1 找到最小值，放到 2 位置上
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i; j < len(arr); j++ { // i ~ N-1 上找到最小值的下标
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {

}
