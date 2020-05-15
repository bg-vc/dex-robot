package task

import (
	"github.com/vincecfl/dex-robot/service/eth"
	"time"
)

func BuyETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		eth.BuyETHHandle()
		time.Sleep(3 * time.Second)
	}
}

func SellETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		eth.SellETHHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		eth.TradeETHHandle()
		time.Sleep(30 * time.Second)
	}
}
