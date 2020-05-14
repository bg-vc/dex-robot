package service

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"time"
)

const (
	RobotTypeKey = "dex:robot:type"
)

func GetDatetime(datetime int64, period string) int64 {
	datetimeStr := time.Unix(int64(datetime), 0).UTC().Format("2006-01-02 15:04:05")
	dt, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.UTC)
	newDatetime := int64(0)
	switch period {
	case "1min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 0, 0, time.UTC).Unix()
	case "5min":
		step := dt.Minute() % 5
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "15min":
		step := dt.Minute() % 15
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "30min":
		step := dt.Minute() % 30
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "1hour":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0, time.UTC).Unix()
	case "4hour":
		step := dt.Hour() % 4
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()-step, 0, 0, 0, time.UTC).Unix()
	case "1day":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.UTC).Unix()
	case "1week":
		step := int(dt.Weekday())
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()-step, 0, 0, 0, 0, time.UTC).Unix()
	case "1mon":
		newDatetime = time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, time.UTC).Unix()
	}
	return newDatetime
}

func GetRobotType() int {
	result := 0
	if !pkg.RedisExists(RobotTypeKey) {
		log.Errorf(nil, "GetRobotType no such key in redis:%v", RobotTypeKey)
		return 0
	}
	val := pkg.GetRedisVal(RobotTypeKey)
	if len(val) == 0 {
		return 0
	}

	if err := json.Unmarshal([]byte(val), &result); err != nil {
		log.Errorf(err, "Unmarshal error")
		return 0
	}
	return result
}

func SetRobotType(robotType int) error {
	if err := pkg.SetRedisVal(RobotTypeKey, robotType, 0); err != nil {
		log.Errorf(err, "SetRedisVal error")
		return err
	}
	return nil
}
