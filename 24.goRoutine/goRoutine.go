package goRoutine

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World!")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Say Hello Function Ending!")
}

func sayHi() {
	fmt.Println("Hiii Prashant!")

}

// if there is more than one go routine then it will run concurrently and work bottom to top like stack (last in first out - LIFO) order of execution
// so of the sayHi() will run first then sayHello() function

func GoRoutine() {
	fmt.Println("Learning Go-Routine")
	go sayHello() // go keyword means it will run concurrently it does not wait for it to finish
	go sayHi()

	time.Sleep(time.Millisecond * 2000)

}
