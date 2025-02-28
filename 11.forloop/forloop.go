package forloop

import "fmt"

func Forloop(table int) {

	var tableArray []int

	for i := 1; i <= 10; i++ {
		tableArray = append(tableArray, table*i)
	}

	fmt.Println(tableArray)

	for index, value := range tableArray {
		fmt.Println(table, "*", index+1, "=", value)

		if index == 4 {
			break
		}
	}

	// in string form
	data := "Hello World"
	for index1, value1 := range data {
		// fmt.Println("Index:", index1, "Value:", value1) // this give u asci value of each character
		fmt.Printf("Index: %d Value: %c \n", index1, value1)
	}

}
