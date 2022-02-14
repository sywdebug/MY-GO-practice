package main

import "fmt"

type Point struct {
	x int
	y int
}
type Point1 struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	point1 := Point1{1, 2}
	a = point
	var b Point
	b = a.(Point)
	fmt.Println(b)
	a = point1
	c, d := a.(Point)
	fmt.Println(c, d)
	fmt.Printf("%T\n",c)
}
