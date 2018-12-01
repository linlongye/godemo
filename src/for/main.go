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

	//打印空心金字塔

	var layers int
	fmt.Println("请输入金字塔层数")
	_, err := fmt.Scanln(&layers)
	/*v1*/
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 1; i <= layers; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
	/*v2*/
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 1; i <= layers; i++ {
			//打印空格
			for k := 1; k <= layers-i; k++ {
				fmt.Print(" ")
			}
			//打印*
			for j := 1; j <= 2*i-1; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	/*v3 空心金字塔*/

	if err != nil {
		fmt.Println(err)
	} else {
		for i := 1; i <= layers; i++ {
			//打印空格
			for k := 1; k <= layers-i; k++ {
				fmt.Print(" ")
			}
			//打印*
			for j := 1; j <= 2*i-1; j++ {
				if i == layers {
					fmt.Print("*")
				} else {
					if j == 1 || j == 2*i-1 {
						fmt.Print("*")
					} else {
						fmt.Print(" ")
					}
				}
			}
			fmt.Println()
		}
	}

	//九九乘法表
	var num2 = 20
	for i := 1; i <= num2; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
