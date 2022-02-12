package model

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

type account struct {
	act     string
	balance float64
	pwd     string
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}

func NewAccount(act string, balance float64, pwd string) *account {
	if len(act) < 6 || len(act) > 10 {
		fmt.Println("账号长度不对")
		return nil
	}
	if balance <= 20 {
		fmt.Println("余额太少了不对")
		return nil
	}
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
		return nil
	}
	return &account{
		act:     act,
		balance: balance,
		pwd:     pwd,
	}
}

func (a *account) SetAct(act string) {
	if len(act) < 6 || len(act) > 10 {
		fmt.Println("账号长度不对")
	} else {
		a.act = act
	}
}

func (a *account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("余额太少了不对")
	}
}
func (a *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密码长度不对")
	} else {
		a.pwd = pwd
	}
}
