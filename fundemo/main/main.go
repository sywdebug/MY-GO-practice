package main

import (
	"fmt"
	"fundemo/utils"
)

func test(n1 int) {
	fmt.Println(n1)
}
func getSumAndSub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
func fbn(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//闭包
func AddUpper() func(int) int {
	var n int = 10
	return func(i int) int {
		n = n + i
		return n
	}
}
func sum2(n1, n2 int) int {
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}
func init() {
	fmt.Println(1212)
}
func main() {
	sum2(10,20)
	fmt.Println(sum(12, 2, 3, 12, 324, 54, 65))
	test(22)
	fmt.Println(fbn(4))
	fmt.Println(getSumAndSub(20, 10))
	fmt.Println(utils.Cal(1, 2))
}
