package main

import "fmt"

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData写的数据=", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData读到的数据=", v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go readData(intChan, exitChan)
	go writeData(intChan)
	for {
		if <-exitChan {
			break
		}
	}
}
