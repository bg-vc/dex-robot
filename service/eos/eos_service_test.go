package eos

import (
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"testing"
)

func TestBuyEOSHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	BuyEOSHandle()
}

func TestSellEOSHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	service.InitSmart()
	SellEOSHandle()
}
