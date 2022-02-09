package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hellojson", func(c *gin.Context) {
		fullpath := "请求路径:" + c.FullPath()
		fmt.Println(fullpath)
		c.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    fullpath,
		})
	})

	engine.GET("/jsonstruct", func(c *gin.Context) {
		fullpath := "请求路径:" + c.FullPath()
		fmt.Println(fullpath)
		resp := Respomse{Code: 1, Message: "OK", Data: fullpath}
		c.JSON(200, &resp)
	})
	engine.Run()
}

type Respomse struct {
	Code    int
	Message string
	Data    interface{}
}
