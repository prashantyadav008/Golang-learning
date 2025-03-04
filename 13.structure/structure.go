package structure

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

func Structure() {
	// var person Person = Person{
	// 	Firstname: "Prashant",
	// 	Lastname:  "Yadav",
	// 	Age:       25,
	// }

	// or

	// var person Person
	// person.Firstname = "Prashant"
	// person.Lastname = "Yadav"
	// person.Age = 25

	var person = new(Person) // new is used to return pointer of struct
	person.Firstname = "Prashant"
	person.Lastname = "Yadav"
	person.Age = 25

	fmt.Println("Person Struct - ", person)
}
