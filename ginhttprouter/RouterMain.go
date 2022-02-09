package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		name := c.Query("name")
		fmt.Println(name)
		c.Writer.Write([]byte("hello," + name))
	})

	engine.POST("/login", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		username, exist1 := c.GetPostForm("username")
		if exist1 {
			fmt.Println("username:" + username)
		}
		password, exist2 := c.GetPostForm("password")
		if exist2 {
			fmt.Println("password:" + password)
		}
		c.Writer.Write([]byte("欢迎" + username + "登录"))
	})

	engine.DELETE("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		fmt.Println(userID)
		c.Writer.Write([]byte("delete用户ID:" + userID))
	})
	engine.Run()
}
