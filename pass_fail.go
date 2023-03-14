package main

import (
	"fmt"
	"hfg-ch4/dates"
	"hfg-ch4/greetings/deutsch"
	"hfg-ch4/keyboard"
	"log"
)

func main() {

	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
	fmt.Println(dates.WeeksToDays(14))
	fmt.Println(dates.DaysPerWeek)

	fmt.Println(dates.DaysToWeeks(9))

	myStr := deutsch.Deutsch()
	fmt.Println(myStr)
}
