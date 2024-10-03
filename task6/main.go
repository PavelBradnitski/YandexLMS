package main

import (
	"fmt"
	"strconv"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}
func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}
func ParseStringToTime(dateString, format string) (time.Time, error) {
	return time.Parse(format, dateString)
}
func TimeAgo(pastTime time.Time) string {
	a := time.Since(pastTime)
	// fmt.Println(a)
	b := strconv.Itoa(int(a.Hours())) + " hours ago"
	return b
}

func NextWorkday(start time.Time) time.Time {

	switch start.Weekday() {
	case time.Friday:
		start = start.AddDate(0, 0, 3)
	case time.Saturday:
		start = start.AddDate(0, 0, 2)
	default:
		start = start.AddDate(0, 0, 1)
	}

	return start
}
func main() {
	start := time.Date(2023, time.October, 6, 0, 0, 0, 0, time.UTC) // A Saturday
	nextWorkday := NextWorkday(start)
	expected := time.Date(2023, time.October, 9, 0, 0, 0, 0, time.UTC) // The following Monday
	fmt.Printf("Expected next workday to be %v, but got %v", expected, nextWorkday)
}
