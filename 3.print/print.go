package print

import (
	"fmt"
)

func Print() {
	age := 25
	name := "Prashant"
	height := 5.8234567

	var emptyString string

	// %d  is for intefer
	// %s is for string
	// %T is for variable types
	// %f is for float
	// %q is for Quoted string

	fmt.Printf("Empty String %q \n", emptyString)

	fmt.Println("age - ", age, "name - ", name, "height - ", height) // println is for format specifier

	fmt.Printf("age is %d and name is %s and height is %.3f \n", age, name, height)

	fmt.Printf("Type of age is %T and name is %T and height is %T \n", age, name, height)

}
