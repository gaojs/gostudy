// 匿名函数的演示
package function

import "fmt"

var funcGlobal = func(a, b int) int {
	return a + b
}

func funcAnonDemo() {
	fmt.Println("匿名函数可以直接调用：")
	res := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println("func(){}(1,2)=", res)
	fmt.Println("匿名函数赋值给函数变量，再调用：")
	f := func(a, b int) int {
		return a + b
	}
	fmt.Println("f=func(){}, f(1,2)=", f(1, 2))
}

func funcAnon() {
	funcAnonDemo()
	fmt.Println("全局匿名函数的演示：")
	fmt.Println("funcGlobal(1,2)=", funcGlobal(1, 2))
}
