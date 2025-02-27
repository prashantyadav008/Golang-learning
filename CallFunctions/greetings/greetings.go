package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

func Hello(name string) string {
	// var message string;
	// message = fmt.Sprintf("Hi, %v. Welcome!", name);
	// 	 or

	// sprintf is only for string
	message := fmt.Sprintf("Hi, %v Welcome!", name)
	return message
}

func Table(number int) []int {

	var result [10]int

	for i := 1; i <= 10; i++ {
		result[i-1] = number * i

	}

	return result[:]

}

func HelloError(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil

}

func RandomNumber() (int, string) {

	number := rand.Intn(10)

	// len("format") ------>> calculate sting length

	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome! \n",
		"Great to see you, %v! \n",
		"Hail, %v! Well met! \n",
	}

	return number, formats[rand.Intn(len(formats))]
}

func GreetOption(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}

	_, greet := RandomNumber()

	message := fmt.Sprintf(greet, name)
	return message, nil
}

func MultiGreets(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := GreetOption(name)
		if err == nil {
			// return nil, err
			messages[name] = message
		}

	}

	return messages, nil

}
