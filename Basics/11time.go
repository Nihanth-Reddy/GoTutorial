package main

import (
	"fmt"
	"time"
)

func main() {
	// Printing local time
	currTime := time.Now()
	fmt.Println("The current time in INDIA is: ", currTime)

	// Converting locatl time to a GMT time
	location, _ := time.LoadLocation("Etc/Greenwich")
	fmt.Println("The current time in GMT is: ", currTime.In(location))

	// String Formatting of time object
	fmt.Printf("The current time in INDIA is: %v and it's type is %T\n", currTime.Format("2006-01-02"), currTime.Format("2006-01-02"))

	// Creating the new time object with custom time of the day
	fmt.Printf("The new time with only date is: %v\n", time.Date(currTime.Year(), currTime.Month(), currTime.Day(), 8, 0, 0, 0, currTime.Location()))

	x := currTime.Truncate(24 * time.Hour)
	fmt.Printf("The current time with only date is: %v\n", x.In(location))
	timeObject, _ := time.Parse("2006-01-02", "2021-04-25")
	fmt.Printf("The time representation of the string 2021-04-25 is: %v\n", timeObject)

	time.Sleep(1 * time.Second)

	greaterTime := time.Now()

	if greaterTime.Before(currTime) {
		fmt.Println("Greater time is great")
	} else {
		fmt.Println("currTime is great")
	}

}

func DateStringParser(utc string) time.Time {
	t, err := time.Parse(time.RFC3339, utc)
	if err != nil {
		panic(err)
	}
	return t
}
