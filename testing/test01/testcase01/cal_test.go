package main

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("addUpper错误，返回值=%v，期望值=%v", res, 55)
	}
	t.Logf("addUpper(10)执行正确")
}
