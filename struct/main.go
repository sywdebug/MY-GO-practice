package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println(cat1)
	cat2 := Cat{
		Name: "小红",
	}
	fmt.Println(cat2)

	var p1 Person
	p1.slice = []int{1, 2, 3}
	fmt.Println(p1)
	p1.slice[0] = 100
	fmt.Println(p1)
	p1.map1 = map[string]string{}
	fmt.Println(p1)
	p1.map1["key"] = "12asd"
	fmt.Println(p1)
	a := 23
	fmt.Println(&a)
	p1.ptr = &a
	fmt.Println(p1)
	fmt.Println()
	fmt.Println()

	var p2 = p1
	fmt.Println(p1)
	fmt.Println(p2)
	p2.map1["key"] = "888"
	p2.map1["key2"] = "66"
	fmt.Println()
	fmt.Println()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println()
	fmt.Printf("p1的地址是%p\n", &p1)
	fmt.Printf("p2的地址是%p\n", &p2)
	fmt.Printf("p1的map1的地址是%p\n", &p1.map1)
	fmt.Printf("p2的map1的地址是%p\n", &p2.map1)
	fmt.Println()
	fmt.Println()
	var p3 = Person{}
	fmt.Println(p3)
	var p4 = new(Person)
	p4.Name = "p4"
	fmt.Println(p4)
	fmt.Println(*p4)
	fmt.Println()
	fmt.Println()
	p5 := &Person{Name: "小红花"}
	p5.Age = 18
	fmt.Println(p5)
	fmt.Println()
	fmt.Println()

	monster := Monster{
		Name:  "老牛",
		Age:   180,
		Skill: "疯子",
	}
	jsonMonster, _ := json.Marshal(monster)
	fmt.Println(string(jsonMonster))

}
