// 接口（多态）
package structure

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct{}

func (p Phone) Start() {
	fmt.Println("手机，开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机，停止工作")
}

type Camera struct{}

func (c Camera) Start() {
	fmt.Println("相机，开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机，停止工作")
}

type Computer struct{}

func (c *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func interfaceUse() {
	computer := Computer{}
	phone := Phone{}
	computer.Working(phone)
	camara := Camera{}
	computer.Working(camara)
}

type I1 interface {
	test1()
}

type I2 interface {
	test2()
}

type I3 interface {
	I1
	I2
	test3()
}

type S1 int

func (s S1) test1() {
	fmt.Println("test1")
}

func (s S1) test2() {
	fmt.Println("test2")
}
func (s S1) test3() {
	fmt.Println("test3")
}

func interfaceExtend() {
	var p *int
	fmt.Printf("p: %T\n", p)
	fmt.Printf("p: %p, %p\n", &p, p)
	if p == nil {
		fmt.Println("p == nil")
	}
	var s S1
	var i1 I1
	fmt.Printf("i1: %T\n", i1)
	fmt.Printf("i1: %p, %p\n", &i1, i1)
	if i1 == nil {
		fmt.Println("i1 == nil")
	}
	i1 = s
	i1.test1()
	fmt.Printf("i1: %T\n", i1)
	var i3 I3 = s
	i3.test3()
	fmt.Printf("i3: %T\n", i3)
}

type INIL interface{}

func nilDemo() {
	var n int = 1
	var i INIL = n
	fmt.Printf("i: %T\n", i)
	fmt.Println("i: ", i)
}

func interfaceDemo() {
	// interfaceUse()
	// interfaceExtend()
	nilDemo()
}
