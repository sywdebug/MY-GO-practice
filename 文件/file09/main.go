package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	filePath := "e:/test3.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件出错，err=", err)
		return
	}
	defer file.Close()
	count := charCount{}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		for _, v := range []rune(str) {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' || v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

		if err == io.EOF {
			break
		}
	}
	fmt.Println(count)
}
