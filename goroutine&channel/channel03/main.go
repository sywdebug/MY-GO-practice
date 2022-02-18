package main

import "fmt"

func main() {
	intChan := make(chan int, 4)
	intChan <- 100
	intChan <- 200
	intChan <- 300
	close(intChan)
	a, a1 := <-intChan
	fmt.Println(a, a1)
	b, b1 := <-intChan
	fmt.Println(b, b1)
	c, c1 := <-intChan
	fmt.Println(c, c1)
	d, d1 := <-intChan
	fmt.Println(d, d1)

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)
	// length := len(intChan2)
	// for i := 0; i < length; i++ {
	// 	fmt.Println(<-intChan2)
	// }
	for v := range intChan2 {
		fmt.Println(v)
	}
}
