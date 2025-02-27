package variables

import "fmt"

var PublicVariable string = "Public Variable" // when 1st letter is capital then it is public and used into other go files as well as package
var privateVariable = "Private Variable"      // when 1st letter is small then it is private and used only in this package

func Variable() {
	// define variable with type as string
	var name string = "Prashant"
	fmt.Println("name - ", name)

	var version = "1.0.0"
	// version = 55 	// not posibile to store other type once variable is defined
	fmt.Println("version - ", version)

	// define variable with type as integer
	var number1 int = 19
	fmt.Println("number1 - ", number1)

	const number2 = 20
	fmt.Println("number2 - ", number2)

	// define variable with type as float
	var dimension float32 = 71.5
	fmt.Println("dimension - ", dimension)

	// define variable with type as boolean
	const isTrue bool = true
	fmt.Println("isTrue - ", isTrue)

	// define variable without type
	person := "Prashant"
	fmt.Println("person - ", person)

	fmt.Println("Public Variable - ", PublicVariable)
	fmt.Println("Private Variable - ", privateVariable)

}
