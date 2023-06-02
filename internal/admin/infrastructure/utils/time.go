package utils

import "time"

const TIME_FORMAT = "2006-01-02 15:04:05"

func FormatTime(t time.Time) string {
	return t.Format(TIME_FORMAT)
}
