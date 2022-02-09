package main

import "fmt"

func maopao(arr *[12]int) {
	fmt.Println(*arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			temp := 0
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(*arr)
}
func main() {
	arr := [...]int{12, 23, 45, 56, 78, 89, 98, 87, 65, 54, 32, 21}
	maopao(&arr)
	fmt.Println(arr)
}
