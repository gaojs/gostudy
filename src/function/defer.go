// defer（延时执行）的演示
package function

import "fmt"

func inc(a int) int {
	defer fmt.Println("1. a=", a)
	a++
	defer fmt.Println("2. a=", a)
	a++
	defer fmt.Println("3. a=", a)
	a++
	return a
}

func funcDefer() {
	fmt.Println("defer的演示：")
	res := inc(1)
	fmt.Println("res=", res)
}
