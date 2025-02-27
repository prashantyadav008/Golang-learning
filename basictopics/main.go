package main

import (
	helloworld "basic-topics/1.helloworld"
	variables "basic-topics/2.variables"
	"fmt"
)

func main() {
	fmt.Println("--------------------- Call Hello World Function ---------------------")
	helloworld.HelloWorld()

	fmt.Println("\n --------------------- Call Variables Function ---------------------")
	variables.Variable()
	fmt.Println("Public Variable used in other package - ", variables.PublicVariable)

}
