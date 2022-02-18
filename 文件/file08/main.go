package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (weitten int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("打开文件出错，err=", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("写文件出错，err=", err)
		return
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	_, err := CopyFile("d:/zjh.jpg", "e:/zjh.jpeg")
	if err != nil {
		fmt.Println("拷贝出错，err=",err)
	} else {
		fmt.Println("拷贝完成")
	}
}
