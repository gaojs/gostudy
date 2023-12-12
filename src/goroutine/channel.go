// 通道Chan
package goroutine

import (
	"fmt"
	"time"
)

func chanDemo1() {
	fmt.Println("通道Chan演示:")
	ch := make(chan int, 1)
	ch <- 10  // 发送数据到通道
	x := <-ch // 从通道接收数据
	close(ch) // 关闭通道
	fmt.Println(x)
}

// 产生N个数，放到ch中
func f1(ch chan int) {
	for i := 0; i < cap(ch); i++ {
		ch <- i
	}
	close(ch)
}

// 从ch1取值，平方后放到ch2中
func f2(ch1, ch2 chan int) {
	for {
		t, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- t * t
	}
	close(ch2)
}

func chanDemo2() {
	fmt.Println("通道Chan演示2:")
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

func worker(id int, jobs, rets chan int) {
	for job := range jobs {
		fmt.Printf("worker: %d, start job: %d\n", id, job)
		time.Sleep(time.Nanosecond)
		rets <- job * job
	}
}

func chanWorker() {
	jobs := make(chan int, 10)
	rets := make(chan int, 10)
	for i := 0; i < cap(jobs); i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 3; j++ {
		go worker(j, jobs, rets)
	}
	for i := 0; i < cap(rets); i++ {
		fmt.Println(<-rets)
	}
}

func chanSelect() {

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		select { // 多路复用，随机选择
		case x := <-ch:
			fmt.Println("取出:", x)
		case ch <- i:
			// fmt.Println("放入:", i)
		default:
			fmt.Println("None")
		}
	}
}

func chanDemo() {
	// chanDemo1()
	// chanDemo2()
	// chanWorker()
	chanSelect()
}
