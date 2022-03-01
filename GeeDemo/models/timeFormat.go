package models

import "time"

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:03:04")
}

func GetDay() string {
	return "20061111111"
}
