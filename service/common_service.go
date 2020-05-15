package service

import (
	"encoding/json"
	"fmt"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"math/rand"
)

const (
	RobotTypeKey = "dex:robot:type:%v"
	BUY          = 1
	SELL         = 2
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

type ResultResp struct {
	Code int           `json:"code"`
	Data PairOrderResp `json:"data"`
}

type PairOrderResp struct {
	Buy   []*PairOrderModel `json:"buy"`
	Sell  []*PairOrderModel `json:"sell"`
	Price float64           `json:"price"`
}

type PairOrderModel struct {
	Price            float64 `json:"price"`
	TotalQuoteAmount float64 `json:"totalQuoteAmount"`
	TotalBaseAmount  float64 `json:"totalBaseAmount"`
	TotalOrder       int64   `json:"totalOrder"`
}

func GetRobotType(pairID int) int {
	key := fmt.Sprintf(RobotTypeKey, pairID)
	result := 0
	if !pkg.RedisExists(key) {
		log.Errorf(nil, "GetRobotType no such key in redis:%v", key)
		return 0
	}
	val := pkg.GetRedisVal(key)
	if len(val) == 0 {
		return 0
	}

	if err := json.Unmarshal([]byte(val), &result); err != nil {
		log.Errorf(err, "Unmarshal error")
		return 0
	}
	return result
}

func SetRobotType(pairID, robotType int) error {
	key := fmt.Sprintf(RobotTypeKey, pairID)
	if err := pkg.SetRedisVal(key, robotType, 0); err != nil {
		log.Errorf(err, "SetRedisVal error")
		return err
	}
	return nil
}

func GetAmount1(price int64) int64 {
	amount1 := int64(0)
	if price <= 1*1e6 {
		amount1 = RandInt64(20, 30) * 1e6
	} else if price > 1*1e6 && price <= 2*1e6 {
		amount1 = RandInt64(10, 15) * 1e6
	} else if price > 2*1e6 {
		amount1 = RandInt64(5, 10) * 1e6
	}
	return amount1
}
