package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	monster := Monster{
		Name:     "牛毛王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.0,
		Skill:    "老牛咆哮",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	fmt.Println(string(data))
}

func testMap() {
	a := map[string]interface{}{
		"name": "红孩儿",
		"age":  30,
	}
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	fmt.Println(string(data))
}

func testSlice() {
	slice := []map[string]interface{}{
		{
			"name":    "jack",
			"age":     7,
			"address": "北京",
		},
		{
			"name":    "tom",
			"age":     17,
			"address": "sasgd京",
		},
	}
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
	}
	fmt.Println(string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	data, _ := json.Marshal(123)
	fmt.Println(string(data))
}
