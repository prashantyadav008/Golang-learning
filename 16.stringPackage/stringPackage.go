package stringPackage

import (
	"fmt"
	"strings"
)

func StringPackage() {
	data := "apple, banana, mango, orange,, pineapple"
	parts := strings.Split(data, ",")
	fmt.Println("Data - ", data)
	fmt.Println("Parts - ", parts, "\t Length of Parts - ", len(parts), parts[3])

	str := "one two three four five two seven eight two ten"
	stringCount := strings.Count(str, "two")
	fmt.Println("String - ", str, "\tString Count - ", stringCount)

	str = "   Hello, Go!    "
	stringTrim := strings.TrimSpace(str)
	fmt.Println("String - ", str, "\t String Trim - ", stringTrim)

	str1 := "Prashant"
	str2 := "Yadav"
	stringConcate := strings.Join([]string{str1, "Kumar", str2}, " ")
	fmt.Println("str1 - ", str1, "\t str2 - ", str2, "\t stringConcate - ", stringConcate)
}
