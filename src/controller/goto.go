package controller

import "fmt"

func gotoDemo() {
	//使用goto语句，跳出循环
	var i = 0
start:
	fmt.Println("start标签")
	for i++; i < 5; i++ {
		if i == 2 {
			fmt.Println("向上跳转，goto start")
			goto start
		} else if i == 4 {
			fmt.Println("向下跳转，goto end")
			goto end
		}
		fmt.Println("循环体，i=", i)
	}
end:
	fmt.Println("end标签")
}
