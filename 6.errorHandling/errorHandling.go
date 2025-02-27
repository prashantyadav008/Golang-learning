package errorHandling

import (
	"fmt"
)

func ErrorHandling() {
	fmt.Println("Sum of two value is: ")

	result, error := divide(10, 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("value cannot divide by zero")
	}

	return a / b, nil
}
