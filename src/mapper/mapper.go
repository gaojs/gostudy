package mapper

import (
	"fmt"
)

func initTest() {
	var s []int
	// s[0] = 3 // panic: index out of range
	// s = append(s, 3) // ok
	println("slice地址：", &s, s)
	fmt.Println("slice内容：", s, len(s), cap(s))
	if s == nil {
		println("slice是nil")
	}
	s = make([]int, 0) // 0无用
	println("slice地址：", &s, s)
	// s[0] = 3 // panic: index out of range
	s = append(s, 3) // ok
	fmt.Println(s)

	var m map[string]int
	// m["gao"] = 0 // panic: assignment to entry in nil map
	println("map地址：", &m, m)
	fmt.Println("map内容：", m, len(m)) // no cap
	if m == nil {
		println("map是nil")
	}
	m = make(map[string]int)
	println("map地址：", &m, m)
	m["gao"] = 0 // ok
	fmt.Println(m)
}

func initDemo() {
	m2 := make(map[string]int, 0)
	fmt.Printf("map类型：%T\n", m2)
	m2["gao"] = 100 // not panic
	m2["hill"] = 80 // not panic
	fmt.Println(m2)
	m3 := map[string]int{"ni": 20}
	fmt.Println(m3)
	// map的值，可以再是map：hill的数学可以考80分
	mm := map[string]map[string]int{"hill": {"math": 80}, "gao": {"art": 20}}
	fmt.Printf("map类型：%T\n", mm)
	fmt.Println(mm)
	// map的键，不可以再是map：数学考80分的，不能是hill
	// mm2 := map[map[string]int]string{{"math": 80}: "hill"}
}

func insertDemo() {
	cities := map[int]string{1: "北京"}
	fmt.Println(cities)
	cities[2] = "上海"
	fmt.Println(cities)
	cities[1] = "杭州"
	fmt.Println(cities)
}

func deleteDemo() {
	cities := map[int]string{1: "北京", 2: "上海", 3: "杭州"}
	fmt.Printf("map类型：%T\n", cities)
	fmt.Println(cities)
	delete(cities, 2) // 有此键，删除有效
	fmt.Println(cities)
	delete(cities, 5) // 无此键，删除无效
	fmt.Println(cities)
	cities = nil // 删除所有元素
	fmt.Println(cities)
	cities = make(map[int]string)
	cities[3] = "苏州"
	fmt.Println(cities)
}

func transDemo() {
	cities := map[int]string{1: "北京", 2: "上海", 3: "杭州"}
	for k, v := range cities {
		fmt.Println(k, v)
	}
	mm := map[string]map[string]int{"hill": {"math": 80}, "gao": {"art": 20}}
	for k, m := range mm {
		for k2, v := range m {
			fmt.Println(k, k2, v)
		}
	}
}

func sliceMap() {
	sm := []map[string]int{}
	fmt.Printf("sm类型：%T\n", sm)
	fmt.Println("cap(sm)=", cap(sm))
	fmt.Println("len(sm)=", len(sm))
	sm = append(sm, map[string]int{"gao": 20})
	sm = append(sm, map[string]int{"hill": 80})
	fmt.Println("cap(sm)=", cap(sm))
	fmt.Println("len(sm)=", len(sm))
	fmt.Println("sm=", sm)
}

func Demo() {
	fmt.Println("map demo")
	initTest()
	// initDemo()
	// insertDemo()
	// deleteDemo()
	// transDemo()
	// sliceMap()
}
