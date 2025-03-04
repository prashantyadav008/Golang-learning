package pointers

import "fmt"

func Pointers() {
	num := 2

	// var ptr *int = &num
	ptr := &num

	fmt.Println("Number:", num, "\t Pointer:", ptr, "\t Pointer:", *ptr)

	var pointer *int
	fmt.Println("Pointer is not assign", pointer) // default value of pointer is nil

	value := 10
	modifyByReference(&value)
	fmt.Println("value --->>> ", value)

}

func modifyByReference(address *int) {
	*address = *address * 2

}
