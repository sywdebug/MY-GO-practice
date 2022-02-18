package main

import (
	"fmt"
	"time"
)

func test2()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("test2() hello,world ", i)
		time.Sleep(time.Second)
	}
}

func test() {
	// go test2()
	for i := 1; i <= 6; i++ {
		fmt.Println("test() hello,world ", i)
		time.Sleep(time.Second)
	}
}
func test3() {
	for i := 1; i <= 6; i++ {
		fmt.Println("test3() hello,world ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程
	go test3()
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,golang ", i)
		time.Sleep(time.Second)
	}
}
