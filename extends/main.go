package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	*Student
}

type Graduate struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}
func main() {
	pupil := Pupil{
		&Student{
			Name: "tom",
			Age:  10,
		},
	}
	pupil.testing()
	pupil.SetScore(90)
	pupil.ShowInfo()
	fmt.Println(pupil.Student)
	fmt.Println(*pupil.Student)
	fmt.Println(pupil.Student.Name)
	fmt.Println(pupil.Name)

	graduate := Graduate{}
	graduate.Name = "marry"
	graduate.Age = 20
	graduate.testing()
	graduate.SetScore(80)
	graduate.ShowInfo()
}
