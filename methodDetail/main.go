package main

import "fmt"

type integer int

type Arr [5]int

func (i integer) print() {
	fmt.Println(i)
}

func (i *integer) change() {
	fmt.Println(i)
	fmt.Println(*i)
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

type qwe struct {
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println(i)

	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)
	fmt.Println(stu.String())
	fmt.Println()
	fmt.Println()
	fmt.Println()
	map1 := map[string]string{}
	struct1 := qwe{}
	slice1 := []int{1, 2, 3, 4, 5}
	int1 := 1
	float1 := 6.4
	bool1 := false
	string1 := "6.4"
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(&map1)
	fmt.Println(&struct1)
	fmt.Println(&slice1)

	fmt.Println(&int1)
	fmt.Println(&float1)
	fmt.Println(&bool1)
	fmt.Println(&string1)
	fmt.Println(&arr)

	arr1:=Arr{1,2,3,4,5}
	arr1.changeArr()
	fmt.Println(arr1)
}
func (arr *Arr)changeArr()  {
	fmt.Println(arr)
	arr[0]=3
}
