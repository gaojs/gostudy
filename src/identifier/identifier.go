// 标识符
package identifier

import (
	"fmt"
)

func iotaDemo() {
	// iota [aɪˈəʊtə] 极少量
	const ( // 遇到const置0
		n1 = iota // 这里是0
		n2 = 100  // 每行增1,这是100
		n3 = iota // 每行增1,这里是2
	)
	fmt.Println(n1, n2, n3)
	const ( // 遇到const置0
		n11 = iota // 这里是0
		n12        // 这里是1
		n13        // 这里是2
		n14 = 10   // 这里是10
		n15        // 这里是10
	)
	fmt.Println(n11, n12, n13, n14, n15)
	const (
		a, b = iota + 1, iota + 2 // iota=0,a=1,b=2
		c, d = iota + 1, iota + 2 // iota=1,c=2,d=3
		e, f = iota + 1, iota + 2 // iota=2,e=3,b=4
	)
	fmt.Println(a, b, c, d, e, f)
	const (
		_  = iota             // 0
		KB = 1 << (10 * iota) // 1
		MB = 1 << (10 * iota) // 2
		GB = 1 << (10 * iota) // 3
		TB = 1 << (10 * iota) // 4
		PB = 1 << (10 * iota) // 5
	)
	fmt.Println(KB, MB, GB, TB, PB)
}

func complexDemo() {
	var score complex64 = complex(1, 2)
	var number complex128 = complex(23.23, 11.11)
	fmt.Print("Score = ", score, "\nNumber = ", number, "\n")
	fmt.Print("Real Score = ", real(score), ", Image Score = ", imag(score), "\n")
	fmt.Print("Real Number = ", real(number), ", Image Number = ", imag(number))
}

func wordDemo() {
	// https://go.dev/ref/spec#Identifiers
	words := "if-else,switch-case-fallthrough-default,for-range-break-continue,var-const-type," +
		"import-package,go-defer,goto-return,func-interface,struct-map,chan-select"
	ids := "any,bool,byte,int-int8-int16-int32-int64,uint-uint8-uint16-uint32-uint64,float32-float64," +
		"rune-string,uintptr,comparable,error,complex64-complex128；true-false,iota,nil；max-min,complex-imag-real," +
		"cap-len-append-copy,new-make-delete-clear,panic-recover,print-println,close"
	fmt.Println("25个保留关键字：", words) // 不能用作常量名或变量名
	fmt.Println("预定义标识符：", ids)     // 基础数据类型,内置函数
	var any int = 1
	fmt.Println(any)  // 不报错,但不推荐,要避免
	var false int = 1 // 不报错,但不推荐,要避免
	fmt.Println(false)
	// complexDemo()
	// iotaDemo()
}

func Demo() {
	wordDemo()
	datatypeDemo()
	operatorDemo()
	builtinDemo()
}
