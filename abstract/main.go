package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	if money <= 0 {
		fmt.Println("金额错误")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}
func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("金额错误")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(pwd string) float64 {
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return 0.0
	}
	return account.Balance
}

func main() {
	qwe := Account{
		AccountNo: "123456789",
		Pwd:       "123456",
		Balance:   10.0,
	}
	fmt.Println("是否插入银行卡，1：是，2：不是")
	in := 0
	fmt.Scanln(&in)
	if in == 1 {
		fmt.Println("请输入密码")
		pwd := ""
		fmt.Scanln(&pwd)
		if pwd == qwe.Pwd {
			for {
				fmt.Println("输入1查询余额，输入2存款，输入3取款，输入0退出")
				var a int
				fmt.Scanln(&a)
				if a == 1 {
					fmt.Printf("您的账号为%v，余额为%v\n", qwe.AccountNo, qwe.Query(pwd))
				} else if a == 2 {
					fmt.Println("请输入存款金额")
					saveMoney := 0.0
					fmt.Scanln(&saveMoney)
					qwe.Deposite(saveMoney, pwd)
					fmt.Printf("您的账号为%v，存储%v，余额为%v\n", qwe.AccountNo, saveMoney, qwe.Query(pwd))
				} else if a == 3 {
					fmt.Println("请输入取款金额")
					taskMoney := 0.0
					fmt.Scanln(&taskMoney)
					qwe.WithDraw(taskMoney, pwd)
					fmt.Printf("您的账号为%v，取出%v，余额为%v\n", qwe.AccountNo, taskMoney, qwe.Query(pwd))
				} else if a == 0 {
					fmt.Println("感谢您的使用，再见")
					break
				}
			}
		} else {
			fmt.Println("密码输入错误")
		}
	} else if in == 2 {
		fmt.Println("不插银行卡怎么操作，八嘎")
	} else {
		fmt.Println("输入指令有误")
	}

}
