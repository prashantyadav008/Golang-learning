package jsonHandling

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FullName string `json:"name"`
	Age      int    `json:"age"`
	IsAdult  bool   `json:"isAdult"`
}

func JsonHandling() {
	person := Person{
		FullName: "Prashant Yadav",
		Age:      25,
		IsAdult:  true,
	}

	fmt.Println("Person Data is - ", person)

	// convert struct into json
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshal-", err)
		return
	}

	fmt.Println("JSON Data is - ", string(jsonData))

	// convert json into struct
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error in Unmarshal-", err)
		return
	}

	fmt.Println("Person Data is - ", personData)
	fmt.Println("Fet FullName in Person Data - ", personData.FullName)
}
