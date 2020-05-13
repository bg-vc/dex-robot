package service

import "time"

func GetDatetime(datetime int64, period string) int64 {
	datetimeStr := time.Unix(int64(datetime), 0).UTC().Format("2006-01-02 15:04:05")
	dt, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.Local)
	newDatetime := int64(0)
	switch period {
	case "1min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 0, 0, time.Local).Unix()
	case "5min":
		step := dt.Minute() % 5
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.Local).Unix()
	case "15min":
		step := dt.Minute() % 15
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.Local).Unix()
	case "30min":
		step := dt.Minute() % 30
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.Local).Unix()
	case "1hour":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0, time.Local).Unix()
	case "4hour":
		step := dt.Hour() % 4
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()-step, 0, 0, 0, time.Local).Unix()
	case "1day":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, time.Local).Unix()
	case "1week":
		step := int(dt.Weekday())
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()-step, 0, 0, 0, 0, time.Local).Unix()
	case "1mon":
		newDatetime = time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, time.UTC).Unix()
	}
	return newDatetime
}
