package main

import (
	"fmt"
	"func/utils"
)

func main() {
	result := utils.Add(10, 20)
	fmt.Println(result)

	res := utils.Calc(12.2, 24, '*')
	fmt.Println(res)
}
