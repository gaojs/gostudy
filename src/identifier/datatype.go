// 数据类型
package identifier

import (
	"fmt"
	"unsafe"
)

func datatypeDemo() {
	var p uintptr // 8字节
	var u uint    // 8字节
	var i int     // 8字节
	var r rune    // 4字节
	var b bool    // 1字节
	fmt.Println("Sizeof(uintptr)=", unsafe.Sizeof(p))
	fmt.Println("Sizeof(uint)=", unsafe.Sizeof(u))
	fmt.Println("Sizeof(int)=", unsafe.Sizeof(i))
	fmt.Println("Sizeof(rune)=", unsafe.Sizeof(r))
	fmt.Println("Sizeof(bool)=", unsafe.Sizeof(b))
	var c64 complex64   // 8字节
	var c128 complex128 // 16字节
	fmt.Println("Sizeof(complex64)=", unsafe.Sizeof(c64))
	fmt.Println("Sizeof(complex128)=", unsafe.Sizeof(c128))
	var s string // 16字节
	var a any    // 16字节
	fmt.Println("Sizeof(string)=", unsafe.Sizeof(s))
	fmt.Println("Sizeof(any)=", unsafe.Sizeof(a))
	var n int = 11
	fmt.Printf("2进制：%b，8进制：%o, 10进制：%d，16进制：%x\n", n, n, n, n)
}
