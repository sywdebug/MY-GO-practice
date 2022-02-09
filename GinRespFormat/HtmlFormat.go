package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//设置html目录
	engine.LoadHTMLGlob("./html/*")
	// 设置静态资源目录
	engine.Static("/img", "./img")
	// GET请求
	engine.GET("/hellohtml", func(c *gin.Context) {
		fullpath := "请求路径:" + c.FullPath()
		fmt.Println(fullpath)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "Gin教程",
			"fullPath": fullpath,
		})
	})
	engine.Run()
}
