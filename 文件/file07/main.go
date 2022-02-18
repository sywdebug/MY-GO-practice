package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path := "e:/test2.txt"
	file2Path := "e:/test.txt"
	content, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("读文件出错，err=", err)
		return
	}
	err = ioutil.WriteFile(file2Path, content, 0666)
	if err != nil {
		fmt.Println("写文件出错，err=", err)
	}
}
