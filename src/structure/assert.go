// 类型断言
package structure

import "fmt"

func typeAssert() {
	var i interface{}
	// var n int = 2
	var n int8 = 2
	i = n
	fmt.Println("i=", i)
	m, ok := i.(int) // 类型断言
	if ok {          // 可以成功
		fmt.Println("m=", m)
	} else {
		fmt.Printf("m=%T\n", m)
	}
}

func assertDemo() {
	typeAssert()
}
