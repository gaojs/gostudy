// 结构体（值类型）
package structure

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDemo() {
	p := Person{"hill", 20}
	fmt.Printf("p类型：%T\n", p)
	// println(p) // illegal types for operand: print
	fmt.Println(p)
	pp := &p
	fmt.Printf("pp类型：%T\n", pp)
	(*pp).Age = 0 // 可以成功修改
	fmt.Println(pp)
	pp.Age = 80 // 也可以成功修改
	fmt.Println(pp)
}

type A struct {
	Num int
}
type A2 A

func typeDemo() {
	var a A = A{1}
	var a2 A2 = A2{0}
	fmt.Println(a2)
	// a2 = a //  cannot use a (variable of type A) as B value in assignment
	a2 = A2(a) // 可以强转，因为A和A2结构体的字段名称和数量一致
	fmt.Println(a2)
}

func jsonDemo() {
	p := Person{"hill", 20}
	bs, err := json.Marshal(p)
	fmt.Println(string(bs), err)
}

func Demo() {
	// initDemo()
	typeDemo()
	jsonDemo()
}
