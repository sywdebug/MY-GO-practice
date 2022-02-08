package main

import "fmt"

func main() {
	var a map[string]string

	a = make(map[string]string)
	a["no1"] = "sywdbeug1"
	a["no2"] = "sywdbeug2"
	a["no3"] = "sywdbeug3"
	a["no1"] = "sywdbeug5"
	a["no4"] = "sywdbeug4"
	a["n4"] = "sywdbeug4"
	a["a4"] = "sywdbeug4"
	fmt.Println(a)

	var b = make(map[string]string)
	b["code"] = "1"
	fmt.Println(b)

	c := map[string]string{
		"code": "1",
		"data": "666",
	}
	fmt.Println(c)

	studentMap := map[string]map[string]string{
		"no1": {
			"name": "二狗",
			"sex":  "男",
		},
		"no2": {
			"name": "铁柱",
			"sex":  "男",
		},
		"no3": {
			"name": "翠花",
			"sex":  "女",
		},
	}
	fmt.Println(studentMap)
	fmt.Println(studentMap["no2"])
	delete(studentMap["no1"], "name")
	fmt.Println(studentMap)

	val, res := studentMap["no1"]
	fmt.Println(val, res)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	for k, v := range studentMap {
		fmt.Println(k, v)
		for k1, v1 := range v {
			fmt.Println(k1, v1)
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(len(studentMap))
	fmt.Println(len(studentMap["no2"]))
}
