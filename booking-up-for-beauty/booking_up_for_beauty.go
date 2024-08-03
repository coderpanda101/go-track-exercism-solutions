package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
		"1/2/2006 15:04:05",             // e.g., "7/25/2019 13:45:00"
		"January 2, 2006 15:04:05",      // e.g., "October 3, 2019 20:32:00"
		"2006-01-02 15:04:05 -0700 MST", // Go's time.Date format
	}

	var parsedTime time.Time
	var err error

	// Try parsing the date string with each layout
	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, date)
		if err == nil {
			return parsedTime
		}
	}

	// If none of the layouts matched, print the error and return a zero time
	fmt.Println("Error parsing date:", err)
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	// parsedDate := Schedule(date)

	// if Schedule(string(time.Now().GoString())).Before(parsedDate) {
	// 	return false
	// }

	return true
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
