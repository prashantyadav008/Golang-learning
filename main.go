package main

import (
	helloworld "basic-topics/1.helloworld"
	variables "basic-topics/2.variables"
	print "basic-topics/3.print"
	function "basic-topics/5.function"
	errorHandling "basic-topics/6.errorHandling"
	"fmt"
)

func main() {
	fmt.Println("--------------------- 1. Call Hello World Function ---------------------")
	helloworld.HelloWorld()

	fmt.Println("\n -------------------- 2. Call Variables Function ---------------------")
	variables.Variable()
	fmt.Println("Public Variable used in other package - ", variables.PublicVariable)

	fmt.Println("\n -------------------- 3. Println vs Printf ---------------------")
	print.Print()

	fmt.Println("\n -------------------- 4. Taking Parameter to User ---------------------")
	// scanf.Scanf()

	fmt.Println("\n -------------------- 5. Functions ---------------------")
	function.Functions(12, 6)

	fmt.Println("\n -------------------- 6. Error Handling ---------------------")
	errorHandling.ErrorHandling()

}
