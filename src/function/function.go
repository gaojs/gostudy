// 函数也是一种类型
package function

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func funcDemo() {
	f := add // f是函数类型的变量
	fmt.Print(" 函数也是一种类型：")
	fmt.Println("f(1, 2)=", f(1, 2))

}

func sum(args ...int) int {
	s := 0 // 求和
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	return s
}

func funcArgs() {
	fmt.Println(" 函数的可变参数的演示：")
	fmt.Println("sum(1)=", sum(1))
	fmt.Println("sum(1, 2)=", sum(1, 2))
	fmt.Println("sum(1, 2, 3)=", sum(1, 2, 3))
}

func Demo() {
	// funcDemo()
	// funcArgs()
	// funcInit()
	// funcAnon()
	funcClos()
}
