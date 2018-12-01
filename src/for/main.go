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
}
