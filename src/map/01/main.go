package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["aa"] = "bb"

	var cities = map[string]string{
		"no4": "cd",
		"no2": "bj",
	}

	for k, v := range cities {
		fmt.Printf("k = %v,v = %v\n", k, v)
	}

	delete(cities, "no2")
	fmt.Println(a["aa"])
	fmt.Println(cities["no4"])

	//初始化map切片
	monster := make([]map[string]string, 2)

	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "100"
	}
	monster = append(monster, map[string]string{"name": "newmonster", "age": "300,"})
	fmt.Println(monster)
	fmt.Println(monster[1] == nil)
}
