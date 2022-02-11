package main

import "fmt"

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=%v gender=%v age=%v id=%v score=%v",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

func main() {
	stu := Student{
		name:   "asd",
		gender: "男",
		age:    12,
		id:     333,
		score:  15.23,
	}
	fmt.Println(stu.say())
	stu2 := Student{
		"asd", "男", 12, 213, 44.6,
	}
	fmt.Println(stu2)

	stu3:=&Student{
		"1235", "男", 12, 213, 44.6,
	}
	fmt.Println(stu3)
}
