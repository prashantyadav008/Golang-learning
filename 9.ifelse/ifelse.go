package ifelse

import "fmt"

func Ifelse(result float64) {

	if result > 90 {
		fmt.Println("Passed!, Got A+")
	} else if result > 80 {
		fmt.Println("Passed!, Got A")
	} else if result > 70 {
		fmt.Println("Passed!, Got B+")
	} else if result > 60 {
		fmt.Println("Passed!, Got B")
	} else if result > 50 {
		fmt.Println("Passed!, Got C+")
	} else if result > 40 {
		fmt.Println("Passed!, Got C")
	} else if result > 32 {
		fmt.Println("Passed!, Got D")
	} else {
		fmt.Println("Failed!, Got F")
	}

	if (result > 32 && result <= 100) || result != 0 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Not in Range")
	}
}
