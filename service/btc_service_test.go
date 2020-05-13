package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
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
	for i:= 0; i < 20; i++ {
		t.Logf("%v", RandInt64(1, 3))
	}
}