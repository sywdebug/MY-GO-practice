package main

import "fmt"

func main() {
	str := "hello哈哈哈sywdbeug"
	slice := str[5:]
	fmt.Println(slice)
	fmt.Println(len(str))
	arr1 := []byte(str)
	fmt.Println(arr1)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	runeArr := []rune(str)
	runeArr[0] = '哈'
	str = string(runeArr)
	fmt.Println(str)
	fmt.Println(fbn(1000))
}
func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}
