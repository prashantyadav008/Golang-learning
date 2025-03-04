package dateTime

import (
	"fmt"
	"time"
)

func DateTime() {
	// By Default format of time is 2006-01-02 15:04:05 or 02-01-2006  (yyyy-mm-dd hh:mm:ss this formate not work in the go language)

	currentTime := time.Now()
	fmt.Println("CurrentTime - ", currentTime)

	formattedDate := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("FormattedDate - ", formattedDate)

	formattedTime := currentTime.Format("02-01-2006 3:04 PM Monday")
	fmt.Println("FormattedTime - ", formattedTime)

	// convert string to date time format
	layoutString := "02/01/2006 03:04 PM"
	dateString := "19/08/1999 03:15 AM"
	formattedData, _ := time.Parse(layoutString, dateString)
	fmt.Println("FormattedData - ", formattedData)

	// increase 1 day of current date
	tomorrow := currentTime.Add(time.Hour * 24)
	fmt.Println("Tomorrow Date - ", tomorrow.Format("02-01-2006"))
}
