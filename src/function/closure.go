// 闭包（函数和环境）的演示
package function

import (
	"fmt"
	"strconv"
)

// 累加器的闭包（等同于类）
func AddUpper() func(int) int {
	var n int = 10 // 等同于类的属性
	var str string = "hello:"
	return func(x int) int {
		n += x // 用到了属性n
		str += strconv.Itoa(x)
		fmt.Println(str)
		return n
	}
}

func funcClos() {
	fmt.Println(" 闭包的演示：")
	f := AddUpper() // 闭包
	for i := 1; i < 4; i++ {
		fmt.Println(f(i))
	}
	f2 := AddUpper() // 新闭包
	for i := 1; i < 4; i++ {
		fmt.Println(f2(i))
	}
}
