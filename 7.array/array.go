package array

import (
	"fmt"
	"slices"
)

func Array() {
	var name [5]string
	var isArray [5]bool

	name[0] = "Item 1"
	name[1] = "Item 2"
	name[2] = "Item 3"
	name[3] = "Item 4"
	name[4] = "Item 5"
	// var table = [8]int{1, 2, 3, 4, 5}
	// or
	var table = [...]int{1, 2, 3, 5: 88, 5, 7} // 5: indicates which index value will be set

	fmt.Println("Name - ", name)
	fmt.Println("isArray - ", isArray)
	fmt.Println("Table Data - ", table, "\t Length of Table - ", len(table))

	var numbers []int
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Numbers - ", numbers)

	var fruits []string
	fruits = append(fruits, "apple")
	fruits = append(fruits, "banana")
	fruits = append(fruits, "mango")
	fruits = append(fruits, "orange")
	fruits = append(fruits, "pineapple")

	fmt.Println("Fruits - ", fruits)
	fmt.Printf("Length %d, \t Type of Fruits Variable - %T \n", len(fruits), fruits)

	// remove data into array using index {first value decide replacable index and 2nd value decide how many index will be covered}
	fruits = append(fruits[:2], fruits[3:]...)
	fmt.Println("After Remove 2nd Element in Fruits - ", fruits)

	// Insert 1 value into array using index
	fruits = slices.Insert(fruits, 2, "grapes")
	fmt.Println("After Insert 1 Element in Fruits - ", fruits)

	// delete last value of fruits array using slice
	fruits = slices.Delete(fruits, 4, 5)
	// or
	// fruits = slices.Delete(fruits, len(fruits)-1, len(fruits))
	fmt.Println("Again Remove 2nd Element in Fruits - ", fruits)

	// copy data into another array
	var newFruits = make([]string, len(fruits))
	copy(newFruits, fruits)
	fmt.Println("Copy Fruits - ", newFruits)

}
