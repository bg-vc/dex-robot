package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
)

func TestBuyEOSHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	BuyEOSHandle()
}

func TestSellEOSHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	SellEOSHandle()
}
