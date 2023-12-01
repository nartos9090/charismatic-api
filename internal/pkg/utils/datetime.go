package utils

import "time"

const FULLDATE_FORMAT_LAYOUT = "2006-01-02"
const DATE_MONTH_FORMAT_LAYOUT = "2 Jan"
const MONTH_YEAR_FORMAT_LAYOUT = "Jan 2006"
const HOURS_IN_DAY = 24

func GetBeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func GetEndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
