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

func typeJudge(args ...interface{}) {

	for i, x := range args {
		switch x.(type) { // 类型断言
		case int, int8, int16, int32, int64:
			fmt.Printf("第%d参数(%v)，是整数\n", i, x)
		case float32, float64:
			fmt.Printf("第%d参数(%v)，是小数\n", i, x)
		case string:
			fmt.Printf("第%d参数(%v)，是字符串\n", i, x)
		default:
			fmt.Printf("第%d参数，是未知类型\n", i)
		}
	}
}

func assertDemo() {
	typeAssert()
	typeJudge(1, 0.1, "hill")
}
