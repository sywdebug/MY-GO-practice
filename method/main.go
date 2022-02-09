package main

import "fmt"

type Person struct {
	Name string
	map1 map[string]string
}

func (person Person) test() {
	fmt.Println(person)
	person.map1["key"] = "yyyyy"
	person.map1["key2"] = "qwe"
}

func (person *Person) test2() {
	fmt.Println("17",person)
	fmt.Println("18",*person)
	person.map1["key"] = "test2"
	person.map1["key2"] = "test22"
}

func (person Person) speak() {
	fmt.Println(person.Name, "是一个好人")
}
func (person Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是", res)
}
func (person Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(person.Name, "计算的结果是", res)
}
func (person Person) getSum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	var person = Person{
		Name: "sywdbeug",
		map1: map[string]string{
			"key": "asd",
		},
	}
	fmt.Println(person)
	person.test()
	person.speak()
	person.jisuan()
	person.jisuan2(10)
	fmt.Println(person.getSum(10, 20))
	person.test()
	fmt.Println(person)
	person.test2()
	fmt.Println(person)
}
