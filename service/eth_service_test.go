package service

import (
	"github.com/vincecfl/dex-robot/pkg"
	"testing"
)

func TestBuyETHHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	BuyETHHandle()
}

func TestSellETHHandle(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	InitSmart()
	SellETHHandle()
}
