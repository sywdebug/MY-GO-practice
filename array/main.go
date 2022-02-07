package main

import "fmt"

func main() {
	var hens [6]float64
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	totalz := 0.0
	for i := 0; i < len(hens); i++ {
		totalz += hens[i]
	}
	fmt.Println(totalz)

	var intArr [3]int
	fmt.Println(intArr)

	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)
	var arr3 = [...]int{1, 2, 3}
	fmt.Println(arr3)
	arr4 := [...]int{1: 200, 2: 12, 3, 7}
	fmt.Println(arr4)
	for _, v := range arr4 {
		fmt.Println(v)
	}
}
