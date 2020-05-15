package task

import (
	"github.com/vincecfl/dex-robot/service/btc"
	"time"
)

func BuyBTCHandle() {
	for {
		btc.BuyBTCHandle()
		time.Sleep(3 * time.Second)
	}
}

func SellBTCHandle() {
	for {
		btc.SellBTCHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeBTCHandle() {
	for {
		btc.TradeBTCHandle()
		time.Sleep(30 * time.Second)
	}
}
