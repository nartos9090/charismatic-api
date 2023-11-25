package utils

import "time"

var FULLDATE_FORMAT_LAYOUT = `2006-01-02`
var DATE_MONTH_FORMAT_LAYOUT = `2 Jan`
var MONTH_YEAR_FORMAT_LAYOUT = `Jan 2006`

func GetBeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func GetEndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
