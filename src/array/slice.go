// 切片（元素个数不固定，引用类型）
package array

import (
	"fmt"
	"unsafe"
)

func capDemo() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	s := a[1:3] // 前闭后开（不包括后限）
	fmt.Println("数组容量（capacity）=", cap(a), "，长度=", len(a))
	fmt.Println("切片容量（capacity）=", cap(s), "，长度=", len(s))
	fmt.Println("数组大小=", unsafe.Sizeof(a)) // 8*6=48=6个元素总大小
	fmt.Println("切片大小=", unsafe.Sizeof(s)) // 8*3=24=1个指针+cap+len
	fmt.Println("数组元素：", a)
	fmt.Println("切片元素：", s)
	s[0] = 20 // 修改切片，数组也变
	fmt.Println("数组元素：", a)
	fmt.Println("切片元素：", s)
}

func makeDemo() {
	s := make([]int, 5, 10) // make(type, len, cap)
	fmt.Println("切片容量=", cap(s), "，长度=", len(s))
	fmt.Println("切片元素：", s) // 整型元素的值默认是0
	s2 := s[3:]             // 可以对切片继续切片
	s2[0] = 1
	fmt.Println("切片元素：", s)
	fmt.Println("切片元素：", s2)
}

func stringDemo() {
	str := "hello杭州"
	fmt.Printf("字符串类型：%T\n", str) // string
	// fmt.Println(cap(str))  // 字符串无cap
	fmt.Println("字符串大小=", unsafe.Sizeof(str)) // 8*2=16=1个指针+len
	s := str[3:]
	fmt.Printf("字符串切片类型：%T\n", s) // string
	// fmt.Println(cap(s))  // 字符串无cap
	fmt.Println("字符串切片大小=", unsafe.Sizeof(s)) // 8*2=16=1个指针+len
	fmt.Println("字符串切片=", s, len(s))
}

func stringChange() {
	str := "hello杭州"
	fmt.Println(str)
	// str[5] = '苏' // 不可修改
	r := []rune(str) // 如尼字母
	fmt.Printf("rune类型：%T\n", r)
	fmt.Println("rune内容：", r)
	r[5] = '苏' // 可修改了
	fmt.Println(string(r))
}

func sliceDemo() {
	// capDemo()
	// makeDemo()
	// stringDemo()
	stringChange()
}
