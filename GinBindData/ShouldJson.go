package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/addstudent", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		var person Person
		if err := c.BindJSON(&person); err != nil {
			log.Fatal()
		}
		fmt.Println("姓名:", person.Name)
		fmt.Println("年龄:", person.Age)
		c.Writer.Write([]byte("添加人员:" + person.Name))
	})
	engine.Run()
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `form:"age"`
}
