package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/_2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January _2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/_2/2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	message := ("You have an appointment on ") + parsedDate.Format("Monday, January 2, 2006, at 15:04.")
	return message
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYearAnniversary := fmt.Sprint(time.Now().Year(), "-09-15 00:00:00 +0000 UTC")
	layout := "2006-01-02 15:04:05 -0700 MST"
	t, _ := time.Parse(layout, currentYearAnniversary)
	return t
}