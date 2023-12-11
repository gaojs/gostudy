// 运算符
package identifier

import "fmt"

func operatorDemo() {
	var i uint8 = 3
	var r1 = ^i       // 按位取反：252=(11111100)2
	var r2 = 0xFF ^ i // 按位异或: 252=(11111100)2
	fmt.Printf("取反：^3=0x%x\n", r1)
	fmt.Printf("异或：0xFF^3=0x%x\n", r2)
}
