package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(RequesInfos())

	engine.GET("/query", func(c *gin.Context) {
		fmt.Println("-----------------中间件间隔")
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  c.FullPath(),
		})
	})
	engine.GET("/hello", func(c *gin.Context) {

	})

	engine.Run()
}

//打印请求信息的中间件
func RequesInfos() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		method := c.Request.Method
		fmt.Println("请求路径：" + path)
		fmt.Println("请求方法：" + method)
		fmt.Println("状态码：", c.Writer.Status())
		c.Next()
		fmt.Println("状态码：", c.Writer.Status())
	}
}
