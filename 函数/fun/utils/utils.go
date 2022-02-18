package utils

import "fmt"

func Cal(n1 float64, n2 float64) float64 {
	var n3 float64
	fmt.Println("输入n3")
	fmt.Scanln(&n3)
	return n1 + n2 + n3
}
