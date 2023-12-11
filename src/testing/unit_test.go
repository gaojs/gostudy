package unit

import "testing"

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	if res != 3 {
		t.Fatalf("用例执行出错！期望值=%v，实际值=%v", 3, res)
	}
	t.Logf("用例执行通过！")
}
