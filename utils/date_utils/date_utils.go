package dateutils

import "time"

const (
	apiDateLayout = "02-01-2006 15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
