package deferKeyword

import "fmt"

func DeferKeyword() {
	// defer use LIFO (last in first out) order
	// defer is used to delay the execution of a function until the surrounding function returns or panics

	fmt.Println("Starting of the Program")         // 1
	defer fmt.Println("Middle of the Program")     // 4
	fmt.Println("End of the Program")              // 2
	defer fmt.Println("Conclusion of the Program") // 3

}
