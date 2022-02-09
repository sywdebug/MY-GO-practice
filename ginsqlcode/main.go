package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库
	// connStr := "root:12345678@tcp(127.0.0.1:3306)/ginsql"
	connStr := "root:0da4fea8c13043c6@tcp(82.156.7.123:7336)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//创建数据表
	//person:id,name,age
	// _, err = db.Exec("create table person(" +
	// 	"id int auto_increment primary key," +
	// 	"name varchar(12) not null," +
	// 	"age int default 1" +
	// 	");")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	return
	// } else {
	// 	fmt.Println("数据库表创建成功!")
	// }

	//插入数据到数据库表
	_, err = db.Exec("insert into person(name,age)"+
		"values(?,?);", "sywdebug2", "23")
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("数据插入成功!")
	}

	//查询数据库
	rows, err := db.Query("select id,name,age from person")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
scan:
	if rows.Next() {
		Person := new(Person)
		err := rows.Scan(&Person.id, &Person.name, &Person.age)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(Person.id, Person.name, Person.age)
		goto scan
	}
}

type Person struct {
	id   int
	name string
	age  int
}
