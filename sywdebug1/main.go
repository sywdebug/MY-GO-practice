package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	connStr := "root:Dh15225217500@tcp(bj-cdb-nnuzvyg4.sql.tencentcdb.com:60179)/sywdebug"
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	engine := gin.Default()
	engine.GET("/checkUser", checkUser)
	engine.POST("/login", login)
	engine.POST("/register", register)
	engine.Run(":3000")
}

//检查用户名是否可用
func checkUser(c *gin.Context) {
	var checkAccount string
	fullpath := "请求路径:" + c.FullPath()
	fmt.Println(fullpath)
	account := c.Query("account")
	sqlStr := "select account from user where account=?"
	row := db.QueryRow(sqlStr, account)
	err := row.Scan(&checkAccount)
	if err != nil {
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"data": "",
			"msg":  "此账号可用",
		})
	} else {
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "此账号不可用",
		})
	}
}

//登陆接口
func login(c *gin.Context) {
	fullPath := "用户登录：" + c.FullPath()
	fmt.Println(fullPath)
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" {
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "账号不能为空!",
		})
	} else if password == "" {
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "密码不能为空!",
		})
	} else {
		sqlStr := "select token,user_id,nickname,role,email,phone from user where account=? and password=?"
		rows, err := db.Query(sqlStr, account, password)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		defer rows.Close()
		fmt.Println(rows.Next())
		user := new(user)
		if rows.Next() {
			err := rows.Scan(&user.token, &user.user_id, &user.nickname, &user.role, &user.email, &user.phone)
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			fmt.Println(user.token, user.user_id, user.nickname, user.role, user.email, user.phone)
			c.JSON(200, map[string]interface{}{
				"code": 1,
				"data": map[string]interface{}{
					"token":    user.token,
					"user_id":  user.user_id,
					"nickname": user.nickname,
					"role":     user.role,
					"email":    user.email,
					"phone":    user.phone,
				},
				"msg": "登陆成功",
			})
		} else {
			c.JSON(200, map[string]interface{}{
				"code": 1,
				"data": "",
				"msg":  "账号或密码填写错误",
			})
		}
	}
}

type user struct {
	token    string
	user_id  string
	nickname string
	role     string
	email    string
	phone    string
}

//注册接口
func register(c *gin.Context) {
	fullpath := "请求路径" + c.FullPath()
	fmt.Println(fullpath)
	token := getToken()
	user_id := getUserId()
	nickname := c.PostForm("nickname")
	account := c.PostForm("account")
	password := c.PostForm("password")
	role := c.PostForm("role")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	if nickname == "" {
		fmt.Println("nickname不能为空")
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "nickname不能为空!",
		})
	} else if account == "" {
		fmt.Println("account不能为空")
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "account不能为空!",
		})
	} else if password == "" {
		fmt.Println("password不能为空")
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "password不能为空!",
		})
	} else if email == "" {
		fmt.Println("email不能为空")
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "email不能为空!",
		})
	} else if phone == "" {
		fmt.Println("phone不能为空")
		c.JSON(200, map[string]interface{}{
			"code": 1000,
			"data": "",
			"msg":  "phone不能为空!",
		})
	} else {
		_, err := db.Exec("insert into user(token,user_id,nickname,account,password,role,email,phone) values(?,?,?,?,?,?,?,?);", token, user_id, nickname, account, password, role, email, phone)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		c.JSON(200, map[string]interface{}{
			"code": 1,
			"data": map[string]interface{}{
				"token":    token,
				"user_id":  user_id,
				"nickname": nickname,
				"account":  account,
				"password": password,
				"role":     role,
				"email":    email,
				"phone":    phone,
			},
			"msg": "注册成功",
		})
	}
}

//获取user_id，使用当前时间戳进行md5加密获得
func getUserId() string {
	timeUnixNano := time.Now().UnixNano()
	str := fmt.Sprintf("%d", timeUnixNano)
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

//获取token，使用当前时间戳进行md5加密再进行md5加密获得
func getToken() string {
	timeUnixNano := time.Now().UnixNano()
	str := fmt.Sprintf("%d", timeUnixNano)
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	data = []byte(md5str1)
	has = md5.Sum(data)
	md5str1 = fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}
