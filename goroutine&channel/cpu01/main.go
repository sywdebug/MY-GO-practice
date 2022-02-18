package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取cpu核数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	//使用cpu
	runtime.GOMAXPROCS(cpuNum / 2)
}
