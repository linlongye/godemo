package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入年龄:")
	_, err := fmt.Scanln(&age)
	if err != nil {
		println(err)
	} else {
		if age > 18 {
			fmt.Println("你的年龄是", age)
		}
	}
}
