package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()")
}
func main() {
	var a AInterface = Stu{}
	a.Say()

	var i integer = 10
	var b AInterface = i
	b.Say()

	var m Monster
	var a1 AInterface=m
	var b1 BInterface=m
	a1.Say()
	b1.Hello()
	m.Hello()
	m.Say()
}
