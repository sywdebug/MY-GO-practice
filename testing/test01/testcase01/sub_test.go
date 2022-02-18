package main

import "testing"

func TestGetSub(t *testing.T) {
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("getSub错误，返回值=%v，期望值=%v", res, 7)
	}
	t.Logf("getSub执行正确")
}
