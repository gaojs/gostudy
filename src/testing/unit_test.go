package unit

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Fatalf("用例执行出错！期望值=%v，实际值=%v", 3, res)
	}
	t.Logf("用例执行通过！")
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}

func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}

func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}
