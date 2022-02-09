package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//GET
	engine.Handle("GET", "/hello", func(c *gin.Context) {
		path := c.FullPath()
		fmt.Println(path)

		name := c.DefaultQuery("name", "hello")
		fmt.Println(name)

		c.Writer.Write([]byte("Hello," + name))
	})
	//POST
	engine.Handle("POST", "login", func(c *gin.Context) {
		fmt.Println(c.FullPath())
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println(username)
		fmt.Println(password)
		c.Writer.Write([]byte(username + "登录"))
	})
	engine.Run()
}
