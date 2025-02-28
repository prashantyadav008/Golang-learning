package function

import "fmt"

func Functions(a, b int) {
	fmt.Println("Sum of two value is: ", Sum(a, b))
	fmt.Println("Multiply of two value is: ", Multiply(a, b))
	fmt.Println("Substract of two value is: ", Substract(a, b))
	fmt.Println("Divide of two value is: ", Divide(a, b))
}

func Sum(a, b int) (result int) {
	result = a + b
	return result
}

func Multiply(a, b int) int {
	return a * b
}

func Substract(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}
