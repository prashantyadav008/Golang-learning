package dataConversion

import (
	"fmt"
	"strconv"
)

func DataConversion() {

	// use case why we use type conversion
	// because when we need to add two values, first is integer and second is float, so it will give error because we can't add integer with float value that why first we need to convert integer to float or vice-versa

	var num int = 10
	fmt.Printf("Number: %d \t Type of Number: %T \n", num, num)

	// type conversion integer into float
	var integerToFloat float64 = float64(num)
	fmt.Printf("Number: %f \t Type of Number: %T \n", integerToFloat, integerToFloat)

	// type conversion integer into string
	var integerToString string = strconv.Itoa(num) //  Itoa ==> I to a
	fmt.Printf("Number: %s \t Type of Number: %T \n", integerToString, integerToString)

	// type conversion string into integer
	var stringToInteger string = "19"
	intData, _ := strconv.Atoi(stringToInteger) //  Atoi ==> A to i
	fmt.Printf("Number: %d \t Type of Number: %T \n", intData, intData)

	// type conversion string into integer
	var stringToFloat string = "154.55543"
	intData1, _ := strconv.ParseFloat(stringToFloat, 64)
	fmt.Printf("Number: %f \t Type of Number: %T \n", intData1, intData1)
}
