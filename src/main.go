package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int  // 8字节
	var u uint // 8字节
	var r rune // 4字节
	var b bool // 1字节
	fmt.Println("int=", unsafe.Sizeof(i))
	fmt.Println("uint=", unsafe.Sizeof(u))
	fmt.Println("rune=", unsafe.Sizeof(r))
	fmt.Println("bool=", unsafe.Sizeof(b))
}
