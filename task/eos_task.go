package task

import (
	"github.com/vincecfl/dex-robot/service/eos"
	"time"
)

func BuyEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		eos.BuyEOSHandle()
		time.Sleep(3 * time.Second)
	}
}

func SellEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		eos.SellEOSHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeEOSHandle() {
	time.Sleep(2 * time.Second)
	for {
		eos.TradeEOSHandle()
		time.Sleep(30 * time.Second)
	}
}
