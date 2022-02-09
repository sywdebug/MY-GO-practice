package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/register", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		var register Register
		if err := c.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)
		c.Writer.Write([]byte(register.UserName + " Register"))
	})
	engine.Run()
}

type Register struct {
	UserName string `form:"name"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}
