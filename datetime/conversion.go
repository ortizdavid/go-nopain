package datetime

import "time"


func StringToDate(strDate string) time.Time {
	date, _ := time.Parse("2006-01-02", strDate)
	return date
}

func StringToDateTime(strDateTime string) time.Time {
	dateTime, _ := time.Parse("2006-01-02 15:04:05", strDateTime)
	return dateTime
}

func DateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func DateTimeToString(date time.Time) string {
	return date.Format("2006-01-02")
}

