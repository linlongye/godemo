package main

import "fmt"

func main() {
outTer:
	for a := 0; a < 10; a++ {
		for b := 0; b < 20; b++ {
			if b == 10 {
				break outTer
			} else {
				fmt.Printf("a的值为%d,b的值为%d\n", a, b)
			}
		}
	}

	// 无限循环
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		} else {
			fmt.Println(sum)
		}
	}
}
