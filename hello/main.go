package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//连接数据库
	connStr := "root:0da4fea8c13043c6@tcp(82.156.7.123:7336)/resources"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//创建数据表
	// person:id,name,age
	// _, err = db.Exec("create table code(" +
	// 	"id int auto_increment primary key," +
	// 	"codeName varchar(12) not null," +
	// 	"fileName varchar(12) not null," +
	// 	"created_at int default 1" +
	// 	");")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("数据库表创建成功!")
	// }
	//插入数据到数据库表
	_, err = db.Exec("insert into style_code(codeName,fileName,filePath,author)"+
		"values(?,?,?,?);", "心跳", "html3", "/path/www", "sywdebug")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("数据插入成功!")
	}

	//查询数据库
	rows, err := db.Query("select * from style_code")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

scan:
	if rows.Next() {
		style_code := new(style_code)
		err := rows.Scan(&style_code.id, &style_code.codeName, &style_code.fileName,&style_code.filePath,&style_code.author,&style_code.created_at)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(style_code.id, style_code.codeName, style_code.fileName,style_code.filePath,style_code.author,style_code.created_at)
		goto scan
	}
}

type style_code struct {
	id         int
	codeName   string
	fileName   string
	filePath   string
	author     string
	created_at string
}
