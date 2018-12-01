package utils

/**
*包中的函数如果是小写字母开头则不能被其他包引用
 */

func Add(a int, b int) int {
	return a + b
}

func Calc(n1 float64, n2 float64, operator byte) float64 {
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
