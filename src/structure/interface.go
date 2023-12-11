// 接口（多态）
package structure

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机，开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机，停止工作")
}

func (p Phone) Call() {
	fmt.Println(p.Name, "手机，可以打电话")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机，开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "相机，停止工作")
}

type Computer struct{}

// 函数参数，是一种多态
func (c *Computer) Working(usb Usb) {
	usb.Start()
	if phone, ok := usb.(Phone); ok {
		phone.Call() // 类型断言成功
	}
	usb.Stop()
}

func interfaceUse() {
	computer := Computer{}
	phone1 := Phone{"小米"}
	phone2 := Phone{"华为"}
	camara := Camera{"佳能"}
	// 数组元素，也是一种多态
	var usbs []Usb = []Usb{phone1, phone2, camara}
	for _, usb := range usbs {
		computer.Working(usb)
	}
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
	interfaceUse()
	// interfaceExtend()
	// nilDemo()
}
