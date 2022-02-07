package main

import "fmt"

func main() {
	var intArr = [...]int{1, 2, 3, 4, 5}
	var slice = intArr[1:4]
	slice[0] = 12
	fmt.Println(intArr, slice)
	fmt.Println("slice的容量", cap(slice))

	var slice2 []float64 = make([]float64, 5, 10)
	slice2[3] = 12
	fmt.Println(slice2, len(slice2), cap(slice2))

	var slice3 = []string{"tom", "jack", "marry"}
	fmt.Println(slice3, len(slice3), cap(slice3))

	for i := 0; i < len(slice3); i++ {
		fmt.Println(slice3[i])
	}

	for _, v := range slice3 {
		fmt.Println(v)
	}
	slice4 := slice[1:3]
	fmt.Println(slice4)
	slice4[0] = 10
	fmt.Println(slice, slice4)

	slice = append(slice, 121, 222, 333)
	slice5 := append(slice, 2121)
	fmt.Println(slice, slice5)
	slice6 := append(slice, slice...)
	fmt.Println(slice6)


}
