package main

import "fmt"

func main() {
	result := add(10, 20)
	fmt.Println(result)

	res := calc(12.2, 24, '*')
	fmt.Println(res)
}

func add(a int, b int) int {
	return a + b
}

func calc(n1 float64, n2 float64, operator byte) float64 {
	var result float64 = 0
	switch operator {
	case '+':
		result = n1 + n2
	case '-':
		result = n1 - n2
	case '*':
		result = n1 * n2
	case '/':
		result = n1 / n2
	}
	return result
}
