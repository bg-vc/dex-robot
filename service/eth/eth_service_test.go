package eth

import (
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"testing"
)

func TestBuyETHHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	BuyETHHandle()
}

func TestSellETHHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	SellETHHandle()
}
