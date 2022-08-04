package calutils

import (
	"time"
)

// Calculates the date of first monday on a given year.
func FirstMondayOfYear(year int) time.Time {
	jan1st := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if jan1st.Weekday() == time.Monday {
		return jan1st
	} else {
		return jan1st.AddDate(0, 0, 8-int(jan1st.Weekday()))
	}
}

func MondayOfTheWeek(year, week int) time.Time {
	firstMonday := FirstMondayOfYear(year)
	return firstMonday.AddDate(0, 0, (week-1)*7)
}
