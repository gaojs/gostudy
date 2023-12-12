// 单元测试
package unit

func Add(a, b int) int {
	return a + b
}

func Fib(n int) int {
	if n < 2 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
