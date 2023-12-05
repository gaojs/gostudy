package main

import "fmt"

func ForDemo() {
	var str = "hello, 北京"
	// fmt.Println(str)
	println(str)
	s := []rune(str) // 转rune
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	println() // 无需转rune
	for _, c := range str {
		fmt.Printf("%c", c)
	}
}
