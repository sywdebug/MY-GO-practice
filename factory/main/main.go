package main

import (
	"fmt"
	"practice/factory/model"
)

func main() {
	// stu := model.student{
	// 	Name:  "sad",
	// 	Score: 22.33,
	// }
	stu := model.NewStudent("tom", 22.33)
	fmt.Println(*stu)
	fmt.Println(stu.Name,stu.GetScore())
}
