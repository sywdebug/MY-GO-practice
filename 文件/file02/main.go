package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("e:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// fmt.Println(reader)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件读取结束")
}
