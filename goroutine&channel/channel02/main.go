package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	// var allChan chan interface{}
	// allChan = make(chan interface{}, 3)
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	cat := Cat{"小猫", 4}
	allChan <- cat

	<-allChan
	<-allChan
	newCat := <-allChan
	fmt.Printf("newCat=%T,newCat=%v\n", newCat, newCat)
	fmt.Println(newCat.(Cat).Name)
}
