package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var f32 float32     // 4字节
	var f64 float64     // 8字节
	var c64 complex64   // 8字节
	var c128 complex128 // 16字节
	fmt.Println("float32=", unsafe.Sizeof(f32))
	fmt.Println("float64=", unsafe.Sizeof(f64))
	fmt.Println("complex64=", unsafe.Sizeof(c64))
	fmt.Println("complex128=", unsafe.Sizeof(c128))
	var i uint8 = 3
	var j = ^i       // 按位取反：252
	var k = 0xFF ^ i // 按位异或:252
	println("j=", j)
	println("k=", k)
}
