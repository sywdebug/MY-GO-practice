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

func (stu Stu) test01() {

}
func (stu Stu) test02() {

}
func (stu Stu) test03() {

}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
	var b BInterface = stu
	b.test01()

	var t interface{}
	t = 8.8
	fmt.Printf("%T\n",t)
	t = "asd"
	fmt.Printf("%T\n",t)
}
