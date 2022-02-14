package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for i, v := range items {
		fmt.Printf("第%v个值是%T类型,值为%v\n", i+1, v, v)
	}
}


func main() {
	n1 := 12
	n2 := 1.2
	n3 := "afgd"
	n4 := true
	TypeJudge(n1,n2,n3,n4)
}
