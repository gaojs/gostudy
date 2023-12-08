// 数组（元素个数固定，数值类型）
package array

import "fmt"

func initDemo() { // 数组的初始化
	var a1 [3]int = [3]int{1, 2, 3} // 数组
	var a2 []int = []int{1, 2, 3}   // 切片
	// var a3 = [3]int{1, 2, 3} // 数组
	var a3 = [...]int{1, 2, 3} // 数组
	var a4 = []int{1, 2, 3}    // 切片

	fmt.Printf("%p, %p\n", &a1, &a1[0]) // 相同
	fmt.Printf("%p, %p\n", &a2, &a2[0]) // 不同
	fmt.Println(a3)
	fmt.Println(a4) // 打印元素
	// println(a3) // 编译不过
	println(a4) // 打印地址

	var names = [3]string{1: "gao", 2: "hill"}
	fmt.Println(names)
}

func transDemo() {
	var a [3]int = [3]int{1, 2, 3} // 数组
	fmt.Println("传统的for(i=...)遍历：")
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	fmt.Println("for-range遍历(index,val)：")
	for index, val := range a {
		fmt.Println(index, val)
	}
}

func Demo() {
	// initDemo()
	// transDemo()
	sliceDemo()
}
