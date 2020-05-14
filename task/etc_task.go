package task

import (
	"github.com/vincecfl/dex-robot/service"
	"time"
)

func BuyETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		service.BuyETHHandle()
		time.Sleep(3 * time.Second)
	}
}

func SellETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		service.SellETHHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeETHHandle() {
	time.Sleep(1 * time.Second)
	for {
		service.TradeETHHandle()
		time.Sleep(30 * time.Second)
	}
}
