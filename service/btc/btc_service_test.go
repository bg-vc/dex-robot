package btc

import (
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"testing"
)

func TestBuyBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	BuyBTCHandle()
}

func TestSellBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	SellBTCHandle()
}

func TestTradeBTCHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	TradeBTCHandle()
}

func TestGenerateRangeNum(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Logf("%v", service.RandInt64(1, 3))
	}
}

func TestGetRobotType(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	t.Logf("%v", service.GetRobotType(1))
}

func TestSetRobotType(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.SetRobotType(1, 2)
}
