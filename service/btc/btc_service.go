package btc

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"github.com/vincecfl/go-common/log"
	"time"
)

const (
	BtcUrl = "https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=1"

	DexContractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	TrxTokenAddr    = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"
	BtcTokenAddr    = "TEQEni8FCPrmdTQPUKAu1DCpm3ZYESjFg8"

	Owner    = "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	OwnerKey = "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07"

	BUY  = 1
	SELL = 2
)

func BuyBTCHandle() {
	result, err := pkg.Get(BtcUrl, false, "")
	if err != nil {
		log.Errorf(err, "BuyBTCHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "BuyBTCHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	price := int64(resp.Data.Price * 1e6)

	if price <= 5*1e5 || (buyLen > 0 && buyList[buyLen-1].Price*1e6 <= 5*1e5) {
		if err := service.SetRobotType(1, 1); err != nil {
			log.Errorf(err, "BuyBTCHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr

	if buyLen >= 20 {
		log.Infof("BuyBTCHandle buyLen more than 20")
		return
	}

	buyPrice := int64(0)

	if buyLen == 0 {
		tempPrice := price
		if tempPrice <= 5*1e5 {
			tempPrice = 6 * 1e5
		}
		// 比当前价格少10000
		buyPrice = tempPrice - 10000
	} else if buyLen < 20 {
		tempPrice := int64(buyList[buyLen-1].Price * 1e6)
		if tempPrice <= 5*1e5 {
			tempPrice = 6 * 1e5
		}
		buyPrice = tempPrice - service.RandInt64(2000, 2500)
	}

	if buyPrice > 0 {
		token2 := TrxTokenAddr
		amount1 := service.GetAmount1(buyPrice)
		amount2 := amount1 * buyPrice / 1e6
		err = service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
		if err != nil {
			log.Errorf(err, "BuyBTCHandle Buy error")
			return
		}
		log.Infof("BuyBTCHandle BuyBTCHandle success")
		return
	}

}

func SellBTCHandle() {
	result, err := pkg.Get(BtcUrl, false, "")
	if err != nil {
		log.Errorf(err, "SellBTCHandle pkg.Get error")
		return
	}
	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "SellBTCHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	sellList := resp.Data.Sell
	sellLen := len(sellList)
	price := int64(resp.Data.Price * 1e6)

	if price >= 30*1e5 || (sellLen > 0 && sellList[sellLen-1].Price*1e6 >= 30*1e5) {
		if err := service.SetRobotType(1, 2); err != nil {
			log.Errorf(err, "SellBTCHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr

	if sellLen >= 20 {
		log.Infof("SellBTCHandle sellLen more than 20")
		return
	}

	sellPrice := int64(0)

	if sellLen == 0 {
		tempPrice := price
		if tempPrice >= 30*1e5 {
			tempPrice = 29 * 1e5
		}
		// 比当前价格多10000
		sellPrice = tempPrice + 10000
	} else if sellLen < 20 {
		tempPrice := int64(sellList[sellLen-1].Price * 1e6)
		if tempPrice >= 30*1e5 {
			tempPrice = 29 * 1e5
		}
		sellPrice = tempPrice + service.RandInt64(2000, 2500)
	}

	if sellPrice > 0 {
		token2 := TrxTokenAddr
		amount1 := service.GetAmount1(sellPrice)
		amount2 := amount1 * sellPrice / 1e6
		err := service.Approve(BtcTokenAddr, userAddr, userKey, DexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "SellBTCHandle Approve error")
			return
		}
		err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
		if err != nil {
			log.Errorf(err, "SellBTCHandle sell error")
			return
		}
		log.Infof("SellBTCHandle SellBTCHandle success")
	}

}

func TradeBTCHandle() {
	result, err := pkg.Get(BtcUrl, false, "")
	if err != nil {
		log.Errorf(err, "TradeBTCHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "TradeBTCHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	sellList := resp.Data.Sell
	sellLen := len(sellList)

	log.Infof("TradeBTCHandle buyLen:%v, sellLen:%v", buyLen, sellLen)

	if buyLen <= 6 || sellLen <= 6 {
		log.Infof("TradeBTCHandle buyLen or sellLen less than 6")
		return
	}

	robotType := service.GetRobotType(1)
	if robotType == 0 {
		log.Errorf(nil, "TradeBTCHandle robotType is 0")
		return
	}

	log.Infof("TradeBTCHandle robotType:%v", robotType)

	currentTime := time.Now().Unix()

	orderType := BUY
	rand := service.RandInt64(1, 101)
	// robotType为2 卖单为主
	if robotType == 2 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			if err := btcSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 5; i++ {
					btcSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := btcBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 2; i++ {
					btcBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := btcBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 5; i++ {
					btcBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := btcSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 2; i++ {
					btcSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		}
		if rand <= 30 {
			orderType = BUY
		} else {
			orderType = SELL
		}
	} else if robotType == 1 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			if err := btcBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 5; i++ {
					btcBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := btcSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 2; i++ {
					btcSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := btcSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 5; i++ {
					btcSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := btcBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 2; i++ {
					btcBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		}
		if rand <= 70 {
			orderType = BUY
		} else {
			orderType = SELL
		}
	}

	if orderType == BUY {
		btcBuy(sellList)
	} else {
		btcSell(buyList)
	}
	return
}

func btcBuy(sellList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "btcBuy error")
		return err
	}
	log.Infof("btcBuy success")
	return nil
}

func btcBuy4Five(sellList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	buyPrice := int64(sellList[index-1].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := int64(0)
	for i := 0; i < index; i++ {
		amount1 += int64(sellList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := service.GetAmount1(buyPrice)
	amount1 += addAmount
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "btcBuy4Five error")
		return err
	}
	log.Infof("btcBuy4Five success")
	return nil
}

func btcSell(buyList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(sellPrice)
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(BtcTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "btcSell Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "btcSell error")
		return err
	}
	log.Infof("btcSell success")
	return nil
}

func btcSell4Five(buyList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	sellPrice := int64(buyList[index-1].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := int64(0)
	for i := 0; i < index; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := service.GetAmount1(sellPrice)
	amount1 += addAmount
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(BtcTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "btcSell4Five Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "btcSell4Five error")
		return err
	}
	log.Infof("btcSell4Five success")
	return nil
}

func btcTrade4Loop(buyList []*service.PairOrderModel, sellList []*service.PairOrderModel, index int) {
	for i := 0; i < index; i++ {
		btcSell(buyList)
		time.Sleep(500 * time.Millisecond)
		btcBuy(sellList)
		time.Sleep(500 * time.Millisecond)
	}
}

func btcBuy4Supply(buyPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "btcBuy4Supply error")
		return err
	}
	log.Infof("btcBuy4Supply success")
	return nil
}

func btcSell4Supply(sellPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := BtcTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(sellPrice)
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(BtcTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "btcSell4Supply Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "btcSell4Supply error")
		return err
	}
	log.Infof("btcSell4Supply success")
	return nil
}
