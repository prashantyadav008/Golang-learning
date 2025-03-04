package main

import (
	helloworld "basic-topics/1.helloworld"
	switchCase "basic-topics/10.switchCase"
	forloop "basic-topics/11.forloop"
	mapping "basic-topics/12.mapping"
	structure "basic-topics/13.structure"
	pointers "basic-topics/14.pointers"
	dataConversion "basic-topics/15.dataConversion"
	stringPackage "basic-topics/16.stringPackage"
	dateTime "basic-topics/17.dateTime"
	deferKeyword "basic-topics/18.deferKeyword"
	fileHandling "basic-topics/19.fileHandling"
	variables "basic-topics/2.variables"
	print "basic-topics/3.print"
	function "basic-topics/5.function"
	errorHandling "basic-topics/6.errorHandling"
	array "basic-topics/7.array"
	slice "basic-topics/8.slice"
	ifelse "basic-topics/9.ifelse"
	"fmt"
)

func main() {
	fmt.Println("---------------------------------------------- 1. Call Hello World Function ----------------------------------------------")
	helloworld.HelloWorld()

	fmt.Println("\n\n---------------------------------------------- 2. Call Variables Function ----------------------------------------------")
	variables.Variable()
	fmt.Println("Public Variable used in other package - ", variables.PublicVariable)

	fmt.Println("\n\n---------------------------------------------- 3. Println vs Printf ----------------------------------------------")
	print.Print()

	fmt.Println("\n\n---------------------------------------------- 4. Taking Parameter to User ----------------------------------------------")
	// scanf.Scanf()

	fmt.Println("\n\n---------------------------------------------- 5. Functions ----------------------------------------------")
	function.Functions(12, 6)

	fmt.Println("\n\n---------------------------------------------- 6. Error Handling ----------------------------------------------")
	errorHandling.ErrorHandling()

	fmt.Println("\n\n---------------------------------------------- 7. Array ----------------------------------------------")
	array.Array()

	fmt.Println("\n\n---------------------------------------------- 8. Slice ----------------------------------------------")
	slice.Slices()

	fmt.Println("\n\n---------------------------------------------- 9. If Else ----------------------------------------------")
	ifelse.Ifelse(83.55)

	fmt.Println("\n\n---------------------------------------------- 10. Switch Cases ----------------------------------------------")
	// month and day
	switchCase.SwitchCase(8, 3)

	fmt.Println("\n\n---------------------------------------------- 11. For Loop ----------------------------------------------")
	// for loop
	forloop.Forloop(5)

	fmt.Println("\n\n---------------------------------------------- 12. Mapping ----------------------------------------------")
	// for loop
	mapping.Mapping()

	fmt.Println("\n\n---------------------------------------------- 13. Structure ----------------------------------------------")
	structure.Structure()

	fmt.Println("\n\n---------------------------------------------- 14. Pointers ----------------------------------------------")
	pointers.Pointers()

	fmt.Println("\n\n---------------------------------------------- 15. Data Conversion ----------------------------------------------")
	dataConversion.DataConversion()

	fmt.Println("\n\n---------------------------------------------- 16. String Package ----------------------------------------------")
	stringPackage.StringPackage()

	fmt.Println("\n\n---------------------------------------------- 17. Date & Time Package ----------------------------------------------")
	dateTime.DateTime()

	fmt.Println("\n\n---------------------------------------------- 18. Defer Keyword ----------------------------------------------")
	deferKeyword.DeferKeyword()

	fmt.Println("\n\n---------------------------------------------- 19. File Handling ----------------------------------------------")
	fileHandling.FileHandling()
}
