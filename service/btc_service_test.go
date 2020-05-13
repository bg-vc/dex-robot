package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
	"time"
)

func TestBuyBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	BuyBTCHandle()
}

func TestSellBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	SellBTCHandle()
}

func TestTradeBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	TradeBTCHandle()
}

func TestGenerateRangeNum(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("%v", RandInt64(1, 3))
	}
}

func TestAAA(t *testing.T) {
	currentTime := time.Now().Unix()

	dateTime00 := GetDatetime(currentTime, "1day")

	dateTime12 := dateTime00 + 60*60*(12-8)

	t.Logf("currentTime:%v, dateTime12:%v, %v", currentTime, dateTime12, currentTime%(30*60))
}
