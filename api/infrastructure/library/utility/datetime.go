package utility

import (
	"strconv"
	"time"
)

// This is not used yet.
func FormatDateTime(argDateTime time.Time) time.Time {
	const format = "2006-01-02 15:04:05"
	datetime := time.Date(
		argDateTime.Year(),
		argDateTime.Month(),
		argDateTime.Day(),
		argDateTime.Hour(),
		argDateTime.Minute(),
		argDateTime.Second(),
		argDateTime.Nanosecond(),
		time.Local,
	)
	return datetime
}

func GetStringDateTimeWithoutHyphen() string {
	now := time.Now()
	return strconv.Itoa(now.Year()) + strconv.Itoa(int(now.Month())) + strconv.Itoa(now.Day())
}
