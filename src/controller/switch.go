package controller

import "fmt"

func switchDemo() {
	print("他是：")
	switch age := 10; {
	case age < 12:
		fmt.Print("儿童；")
		fallthrough
	case age < 18:
		fmt.Print("少年；")
	case age > 60:
		fmt.Print("老人；")
	default:
		fmt.Print("成年；")
	}
	println()
}
func switchType() {
	// var x interface{} = 10
	var x any = nil
	fmt.Print("x类型：")
	switch t := x.(type) {
	case int, uint:
		fmt.Println("整数")
	case float32, float64:
		fmt.Println("小数")
	case string:
		fmt.Println("字符串")
	default:
		fmt.Printf("%T\n", t)
	}
}
