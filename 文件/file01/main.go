package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("E:/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Println(file)

	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
