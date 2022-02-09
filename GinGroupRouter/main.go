package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	routerGroup := engine.Group("/user")
	routerGroup.POST("/register", registerHandle)
	routerGroup.GET("/login", loginHandle)
	routerGroup.DELETE("/:id", deleteHandle)
	engine.Run()
}

func registerHandle(c *gin.Context) {
	fullPath := "用户注册：" + c.FullPath()
	fmt.Println(fullPath)
	c.Writer.WriteString(fullPath)
}
func loginHandle(c *gin.Context) {
	fullPath := "用户登录：" + c.FullPath()
	fmt.Println(fullPath)
	c.Writer.WriteString(fullPath)
}
func deleteHandle(c *gin.Context) {
	fullPath := "用户删除：" + c.FullPath()
	userID := c.Param("id")
	fmt.Println(fullPath + " " + userID)
	c.Writer.WriteString(fullPath + " " + userID)
}
