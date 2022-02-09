package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellobyte", func(c *gin.Context) {
		fullPath := "请求路径:" + c.FullPath()
		fmt.Println(fullPath)
		c.Writer.Write([]byte(fullPath))
	})

	engine.GET("/hellostring", func(c *gin.Context) {
		fullPath := "请求路径:" + c.FullPath()
		fmt.Println(fullPath)
		c.Writer.WriteString(fullPath)
	})

	engine.Run()
}
