package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入年龄")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println(err)
	} else {
		// go中的case语句块中不需要添加break，编译器会自动添加
		switch age {
		case 1, 2, 3, 4:
			fmt.Println("you are too young")
		case 20:
			fmt.Println("you are too old")
			//default语句不是必须的
		default:
			fmt.Println("you are ok")
		}
	}
}
