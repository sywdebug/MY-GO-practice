package main

import "fmt"

func main() {

	map1 := map[int]int{
		10: 100,
		1:  2,
		3:  4,
		5:  6,
	}
	fmt.Println(map1)

	monsters := []map[string]string{
		{"name": "狐狸精", "age": "200"},
		{"name": "孙猴子", "age": "1000"},
		{"name": "牛魔王", "age": "2000"},
	}
	monsters2 := map[string]string{
		"name": "猪八戒", "age": "500",
	}
	fmt.Println(monsters)
	monsters = append(monsters, monsters2)
	fmt.Println(monsters)
}
