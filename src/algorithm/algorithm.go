// 算法（排序和查找）
package algorithm

import "fmt"

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j-1] > a[j] { // 交换
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}
}

func binarySearch(a []int, e int) int {
	left := 0
	right := len(a) - 1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == e {
			return mid
		} else if a[mid] > e {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func bsDemo() {
	a := []int{8, 98, 3, 20, 123}
	fmt.Println("原数组：", a)
	bubbleSort(a)
	fmt.Println("排序后：", a)
	fmt.Println("1的位置：", binarySearch(a, 1))
	fmt.Println("8的位置：", binarySearch(a, 8))
	fmt.Println("50的位置：", binarySearch(a, 50))
	fmt.Println("98的位置：", binarySearch(a, 98))
	fmt.Println("125的位置：", binarySearch(a, 125))
}

func Demo() {
	// bsDemo()
	sortDemo()
}
