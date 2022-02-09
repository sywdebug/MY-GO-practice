package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.FullPath())

		var student Student
		err := c.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		c.Writer.Write(([]byte("hello," + student.Name)))
	})
	engine.Run()
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
