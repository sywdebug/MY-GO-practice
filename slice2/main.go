package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 10)
	slice[0]=88
	copy(slice2, slice)
	fmt.Println(slice)
	fmt.Println(slice2)
}