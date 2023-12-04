package main

import "fmt"

func switchDemo() {
	print("他是：")
	switch age := 10; {
	case age < 12:
		println("儿童")
		fallthrough
	case age < 18:
		println("少年")
	case age > 60:
		println("老人")
	default:
		println("成年")
	}
}
func switchType() {
	// var x interface{} = 10
	var x any = nil
	print("x类型：")
	switch t := x.(type) {
	case int, uint:
		println("整数")
	case float32, float64:
		println("小数")
	case string:
		println("字符串")
	default:
		fmt.Printf("%T\n", t)
	}
}

func forDemo() {
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

func main() {
	// switchDemo()
	// switchType()
	forDemo()
}
