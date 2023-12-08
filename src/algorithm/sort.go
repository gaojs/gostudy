package algorithm

import (
	"fmt"
	"math/rand"
	"sort"
)

func sortInt() {
	a := []int{8, 90, 32, 3}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)
}

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (x HeroSlice) Len() int {
	return len(x)
}

func (x HeroSlice) Less(i, j int) bool {
	if x[i].Age != x[j].Age {
		return x[i].Age < x[j].Age
	} else {
		return x[i].Name < x[j].Name
	}
}

func (x HeroSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func sortHero() {
	var hs HeroSlice
	for i := 0; i < 5; i++ {
		id := rand.Intn(100)
		age := rand.Intn(100)
		name := fmt.Sprintf("英雄~%d", id)
		h := Hero{Name: name, Age: age}
		hs = append(hs, h)
	}
	fmt.Println(hs)
	sort.Sort(hs)
	fmt.Println(hs)
}

func sortDemo() {
	// sortInt()
	sortHero()
}
