package main

import (
	"container/list"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "root:0da4fea8c13043c6@tcp(82.156.7.123:7336)/resources"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	engine := gin.Default()
	engine.GET("/show", func(c *gin.Context) {
		fullpath := "请求路径:" + c.FullPath()
		fmt.Println(fullpath)
		rows, err := db.Query("select * from style_code;")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		codeData := list.New()
		fmt.Println(codeData)
		// var codeDataItem map[string]string
		codeDataItem := make(map[string]string)
	scan:
		if rows.Next() {
			code := new(code)
			err := rows.Scan(&code.id, &code.codeName, &code.fileName, &code.filePath, &code.author, &code.createTime)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
			codeDataItem["id"] = code.id
			codeDataItem["codeName"] = code.codeName
			codeDataItem["fileName"] = code.fileName
			codeDataItem["filePath"] = code.filePath
			codeDataItem["author"] = code.author
			codeDataItem["createTime"] = code.createTime
			fmt.Println(code.id, code.codeName, code.fileName, code.filePath, code.author, code.createTime)
			fmt.Println(codeDataItem)
			codeData.PushFront(codeDataItem)
			fmt.Println(codeData)
			goto scan
		}

		fmt.Println("查询成功", codeData)
		c.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "数据查询成功",
			"data":    codeData,
		})
	})
	engine.POST("/insert", func(c *gin.Context) {
		fullpath := "请求路径:" + c.FullPath()
		fmt.Println(fullpath)
		codeName := c.PostForm("codeName")
		fileName := c.PostForm("fileName")
		filePath := c.PostForm("filePath")
		author := c.PostForm("author")
		fmt.Println(codeName)
		fmt.Println(fileName)
		fmt.Println(filePath)
		fmt.Println(author)
		if codeName == "" || fileName == "" || filePath == "" || author == "" {
			fmt.Println("不能为空")
			c.JSON(200, map[string]interface{}{
				"code":    2,
				"message": "codeName和fileName和filePath和author均不可为空",
				"data":    fullpath,
			})
		} else {
			_, err = db.Exec("insert into style_code(codeName,fileName,filePath,author)"+
				"values(?,?,?,?);", codeName, fileName, filePath, author)
			if err != nil {
				log.Fatal(err.Error())
				return
			} else {
				fmt.Println("数据插入成功")
				c.JSON(200, map[string]interface{}{
					"code":    1,
					"message": "数据插入成功",
					"data":    fullpath,
				})
			}
		}
	})

	engine.Run(":8000")
}

type code struct {
	id         string
	codeName   string
	fileName   string
	filePath   string
	author     string
	createTime string
}
