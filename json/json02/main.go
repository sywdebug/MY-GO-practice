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

func unmarshalStruct() {
	str := "{\"name\":\"牛毛王\",\"age\":500,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"老牛咆哮\"}"
	monster := Monster{}
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
	}
	fmt.Println(monster)

}

func unmarshalMap() {
	str := "{\"age\":30,\"name\":\"红孩儿\"}"
	a := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
	}
	fmt.Println(a)
}

func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"},{\"address\":\"sasgd京\",\"age\":17,\"name\":\"tom\"}]"
	slice := []map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败，err=", err)
	}
	fmt.Println(slice)

}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
