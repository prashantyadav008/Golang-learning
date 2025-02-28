package slice

import (
	"fmt"
)

func Slices() {

	var numbers []int
	numbers = append(numbers, 11, 22, 33, 44, 55, 66, 77)

	fmt.Println("Numbers, Get Length & Get Capacity - ", numbers, len(numbers), cap(numbers))

	numbersUsingMake := make([]int, 3, 5) // if capacity will be set then it will be double other wise it will be same as length
	fmt.Println("Numbers Using Make, Get Length & Get Capacity - ", numbersUsingMake, len(numbersUsingMake), cap(numbersUsingMake))

	// after put some values into array the capacity will be double
	numbersUsingMake = append(numbersUsingMake, 11, 22, 33)
	fmt.Println("Numbers Using Make, Get Length & Get Capacity - ", numbersUsingMake, len(numbersUsingMake), cap(numbersUsingMake))

	// after put some values into array the capacity will be double
	numbersUsingMake = append(numbersUsingMake, 11, 22, 33, 44, 55, 66, 77, 88)
	fmt.Println("Numbers Using Make, Get Length & Get Capacity - ", numbersUsingMake, len(numbersUsingMake), cap(numbersUsingMake))

}
