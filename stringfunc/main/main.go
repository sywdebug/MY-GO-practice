package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//字符串转整数
	n, err := strconv.Atoi("aa")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}
	//整数转字符串
	str := strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", str, str)
	//字符串转byte
	var bytes = []byte("hello world")
	fmt.Printf("bytes=%v\n", bytes)
	//byte转字符串
	str = string([]byte{98, 99, 100})
	fmt.Println("str=", str)

	fmt.Println(strings.Contains("qweasd", "eas"))
	fmt.Println(strings.Count("qweasdqwasqw", "q"))
	fmt.Println(strings.EqualFold("asd", "Asd"))
	fmt.Println(strings.Index("asdfghd", "d"))
	fmt.Println(strings.LastIndex("asdfghd", "d"))
	fmt.Println(strings.Replace("asdfghdedgtgd", "d","第",-1))
	fmt.Println(strings.Split("hello world haha,ss"," "))
	fmt.Println(strings.ToLower("heLlo wOrld hAha,ss"))
	fmt.Println(strings.ToUpper("heLlo wOrld hAha,ss"))
	fmt.Println(strings.TrimSpace("  heLlo wOrld hAha,ss    "))
	fmt.Println(strings.Trim("sheLlo wOrld hAha,ss","hs"))
	fmt.Println(strings.HasPrefix("https://sywdebug.com","http"))
	fmt.Println(strings.HasSuffix("https://sywdebug.com","com"))
}
