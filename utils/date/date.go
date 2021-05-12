package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	DBDateLayout  = "2006-01-02 15:04:05"
)

// GetNow function returns current UTC time object
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString function returns current UTC time in string format
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(DBDateLayout)
}
