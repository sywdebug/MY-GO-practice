package main

import "fmt"

func main() {
	arr := [...]string{"鼠", "牛", "虎", "兔", "龙"}
	var s = ""
	fmt.Println("请输入查找的信息")
	fmt.Scan(&s)
	for i, v := range arr {
		if v == s {
			fmt.Printf("找到了，在第%d个，是%v\n", i+1, v)
		}
	}

	index := -1
	for i, v := range arr {
		if v == s {
			index = i
		}
	}
	if index != -1 {
		fmt.Printf("找到了，在第%d个，是%v\n", index+1, arr[index])
	} else {
		fmt.Println("没找到")
	}
}
