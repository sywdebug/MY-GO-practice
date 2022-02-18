package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.OpenFile("e:/test2.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	file, err := os.OpenFile("e:/test2.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	str := "hello,司帅\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
