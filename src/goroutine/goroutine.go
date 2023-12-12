// GO协程
package goroutine

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("hello", i)
	wg.Done() // 完成+1
}
func goDemo() {
	wg.Add(100) //等待+1
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	fmt.Println("main")
	wg.Wait()

}
func Demo() {
	// goDemo()
	// chanDemo()
	lockDemo()
}
