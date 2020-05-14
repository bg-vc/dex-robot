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
	for i := 0; i < 20; i++ {
		t.Logf("%v", RandInt64(1, 3))
	}
}

func TestGetRobotType(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	t.Logf("%v", GetRobotType(1))
}

func TestSetRobotType(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	SetRobotType(1, 2)
}
