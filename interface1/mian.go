package main

import "fmt"

type AInterface interface {
	Say(name string)
}

type BInterface interface {
	Hello()
}

type Stu struct {
	Name string
}

func (stu Stu) Say(name string) {
	stu.Name = name
	fmt.Println(stu.Name)
}
func (stu Stu) Hello() {
	fmt.Println("hello")
}

func main() {
	stu := Stu{}
	var a AInterface = stu
	a.Say("sywdebug")
	var b BInterface = stu
	b.Hello()
}
