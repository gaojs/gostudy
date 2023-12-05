package controller

import "fmt"

func forDemo() {
	var str = "hello, 北京"
	// fmt.Println(str)
	fmt.Println("原始字符串：", str)
	fmt.Print("for i=... 要转rune：")
	s := []rune(str) // 转rune
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	println()
	// 无需转rune
	fmt.Print("for-range 无需转rune：")
	for _, c := range str {
		fmt.Printf("%c", c)
	}
	println()
}
