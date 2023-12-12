// 并发与锁
package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int
var wt sync.WaitGroup

// var lock sync.Mutex // 互斥锁
var lock sync.RWMutex // 读写锁

func inc() {
	for i := 0; i < 1000; i++ {
		x += 1
	}
	wt.Done()
}

func incLock() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wt.Done()
}

func mutexDemo() {
	n := 10
	wt.Add(n)
	for i := 0; i < n; i++ {
		go incLock()
		// go inc()
	}
	wt.Wait()
	fmt.Println("执行结束:", x)
}
func one() {
	fmt.Println("仅执行一次! Once!")
}
func onceDemo() {
	n := 10
	wt.Add(n)
	var once sync.Once
	for i := 0; i < n; i++ {
		go func() {
			// one() // 并发
			once.Do(one)
			wt.Done()
		}()
	}
	wt.Wait()
}

func atomDemo() {
	atomic.AddInt32()
}

func lockDemo() {
	mutexDemo()
	onceDemo()
	atomDemo()
}
