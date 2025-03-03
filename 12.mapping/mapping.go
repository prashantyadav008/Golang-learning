package mapping

import "fmt"

func Mapping() {
	studentGrades := make(map[string]int)

	studentGrades["John"] = 90
	studentGrades["Mark"] = 85
	studentGrades["Alice"] = 76
	studentGrades["Robert"] = 70

	fmt.Println("Student Grades - ", studentGrades, "\t Mark's Grade - ", studentGrades["Mark"])

	studentGrades["Mark"] = 55
	fmt.Println("New Mark's Grade - ", studentGrades["Mark"])

	// move key-value pair
	delete(studentGrades, "Mark")
	fmt.Println("Student Grades - ", studentGrades)

	// check any key exist or not
	grades, exists := studentGrades["Mark"]
	fmt.Println("Mark's Exits or not - ", exists, "\t Grades - ", grades)
	grades1, exists1 := studentGrades["Prashant"]
	fmt.Println("Prashant Exits or not - ", exists1, "\t Grades - ", grades1)
	grades2, exists2 := studentGrades["John"]
	fmt.Println("John Exits or not - ", exists2, "\t Grades - ", grades2)

	person := map[string]int{
		"John":   90,
		"Mark":   85,
		"Alice":  76,
		"Robert": 70,
	}

	fmt.Println("Person Data - ", person)

}
