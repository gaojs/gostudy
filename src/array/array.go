package array

import "fmt"

func initDemo() { // 数组的初始化
	var a1 [3]int = [3]int{1, 2, 3} // 数组
	var a2 []int = []int{1, 2, 3}   // 切片
	// var a3 = [3]int{1, 2, 3} // 数组
	var a3 = [...]int{1, 2, 3} // 数组
	var a4 = []int{1, 2, 3}    // 切片

	fmt.Println(&a1, a1)
	fmt.Println(&a2, a2)
	fmt.Println(&a3, a3)
	fmt.Println(&a4, a4)
	fmt.Printf("%p, %p\n", &a2, a2)
	println(&a2, a2)

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
	initDemo()
	transDemo()
}
