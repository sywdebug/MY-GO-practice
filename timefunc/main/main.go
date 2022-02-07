package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 1. 获取当前时间
	fmt.Println(now)
	// 获取年月日时分秒
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 格式化日期和时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	// 休眠
	// for i := 0; ; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100)
	// 	if i == 100 {
	// 		break
	// 	}
	// }
	// unix和unixNano
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}
