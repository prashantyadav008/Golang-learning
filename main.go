package main

import (
	helloworld "basic-topics/1.helloworld"
	variables "basic-topics/2.variables"
	print "basic-topics/3.print"
	scanf "basic-topics/4.scanf"
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
	scanf.Scanf()
}
