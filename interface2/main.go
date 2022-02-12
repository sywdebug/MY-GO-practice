package main

import "fmt"

type BInterface interface {
	test01()
}
type CInterface interface {
	test02()
}
type AInterface interface {
	BInterface
	CInterface
	test03()
}
type Stu struct {
}

func (s Stu) test01() {

}
func (s Stu) test02() {

}
func (s Stu) test03() {

}
func main() {
	var s Stu
	var a AInterface = s
	a.test01()
	var i interface{}
	fmt.Println(i)
	i = "asd"
	fmt.Println(i)
	i = 182
	fmt.Println(i)
	i = 123.4524
	fmt.Println(i)
	i = []int{1, 2, 3, 4}
	fmt.Println(i)
	i = false
	fmt.Println(i)
}
