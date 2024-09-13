package main

import (
	"fmt"
	"time"
)

func main() {

	// yyyy = 2006, mm = 01, dd = 02
	// 2006-01-02 15:04:05

	var currentTime time.Time = time.Now()
	fmt.Println("Current time:", currentTime)

	formatDate := currentTime.Format("02-01-2006")
	fmt.Println("Formatted time:", formatDate)

	formatDate = currentTime.Format("2006/01/02")
	fmt.Println("Formatted time:", formatDate)

	formatDate = currentTime.Format("02-01-2006, Monday")
	fmt.Println("Formatted time:", formatDate)

	formatDate = currentTime.Format("02-01-2006, 15:04:05")
	fmt.Println("Formatted time:", formatDate)

	formatDate = currentTime.Format("02-01-2006, 3:04 PM")
	fmt.Println("Formatted time:", formatDate)

	dateStr := "2023-11-25" // same format as string date
	layoutStr := "2006-01-02"
	strToDate, _ := time.Parse(layoutStr, dateStr)
	fmt.Println("String to date:", strToDate)

	// add 1 more day in current time
	newDate := currentTime.Add(24 * time.Hour)
	fmt.Println("new date:", newDate)
	fmt.Println("Formatted new date:", newDate.Format("2006/01/02 Monday"))

}
