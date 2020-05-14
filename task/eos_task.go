package task

import (
	"github.com/vincecfl/dex-robot/service"
	"time"
)

func BuyEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		service.BuyEOSHandle()
		time.Sleep(3 * time.Second)
	}
}

func SellEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		service.SellEOSHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		service.TradeEOSHandle()
		time.Sleep(30 * time.Second)
	}
}
