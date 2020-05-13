package task

import (
	"github.com/vincecfl/dex-robot/service"
	"time"
)

func BuyBTCHandle() {
	for {
		service.BuyBTCHandle()
		time.Sleep(3 * time.Second)
	}
}


func SellBTCHandle() {
	for {
		service.SellBTCHandle()
		time.Sleep(3 * time.Second)
	}
}

func TradeBTCHandle() {
	for {
		service.TradeBTCHandle()
		time.Sleep(30 * time.Second)
	}
}