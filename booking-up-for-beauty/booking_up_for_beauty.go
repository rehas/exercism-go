package booking

import (
    "time"
    "fmt"
    )

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
    longForm := `1/2/2006 15:04:05`
	t,_ := time.Parse(longForm, date)
    return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	longForm := `January 2, 2006 15:04:05`
	t,_ := time.Parse(longForm, date)
    return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	longForm := `Monday, January 2, 2006 15:04:05`
	t,_ := time.Parse(longForm, date)
    return t.Hour() >= 12 && t.Hour() <=18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	longForm := `1/2/2006 15:04:05`
	t,_ := time.Parse(longForm, date)
    h,m,_ := t.Clock()
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), h, m)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
    return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
