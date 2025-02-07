package main

import (
	"fmt"
	"time"

	"log"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	fmt.Println("Hello World!")
	fmt.Println(time.Now())
	fmt.Println(quote.Go())

	// print greetings hello function
	{
		message := greetings.Hello("Prashant")
		fmt.Println(message)
	}

	// print greeting table function
	{
		table := greetings.Table(5)
		fmt.Println("Table of 5 is", table)
	}

	// print greetings hello error function
	{
		message, err := greetings.HelloError("")

		// print logs
		// log.Printf("message --->>> %v \t  error --_>> %v", message, err)

		if err != nil {
			// print logs
			log.Printf("message --->>> %v \t  error --_>> %v", message, err)

			// // exit code working like die
			// log.Fatal(err)
		}

		fmt.Println(message)
	}

	// print random number function
	{
		message, error := greetings.GreetOption("Prashant")
		if error != nil {
			log.Printf("error --_>> %v", error)
		}
		fmt.Println(message)
	}

	// print multi greeting
	{
		message, error := greetings.MultiGreets([]string{"Prashant", "", "Prashant 1", "Prashant 2"})
		if error != nil {
			log.Printf("error --_>> %v", error)
		}
		fmt.Println(message)
	}
}
