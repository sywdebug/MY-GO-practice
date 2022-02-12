package main

import (
	"fmt"
	"practice/encapsulate/model"
)

func main() {
	p := model.NewPerson("tom")
	fmt.Println(p)
	p.SetAge(18)
	p.SetSal(8000)
	fmt.Println(p)
	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(p.GetSal())

	a := model.NewAccount(
		"tom", 666, "123456",
	)
	if a != nil {
		fmt.Println(123)
		a.SetAct("sywdebug")
		a.SetBalance(9999999)
		// a.SetPwd("123456")
		fmt.Println(a)
	} else {
		fmt.Println(666)
	}
}
