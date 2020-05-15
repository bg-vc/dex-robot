package eth

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"github.com/vincecfl/go-common/log"
	"time"
)

const (
	EthUrl = "https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=2"

	DexContractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	TrxTokenAddr    = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"
	EthTokenAddr    = "TYLNNGZib5fH77Xw3VaWx5J89RiSVfWfbL"

	Owner    = "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	OwnerKey = "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07"

	BUY  = 1
	SELL = 2
)

func BuyETHHandle() {
	result, err := pkg.Get(EthUrl, false, "")
	if err != nil {
		log.Errorf(err, "BuyETHHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "BuyETHHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	price := int64(resp.Data.Price * 1e6)

	if price <= 5*1e5 || (buyLen > 0 && buyList[buyLen-1].Price*1e6 <= 5*1e5) {
		if err := service.SetRobotType(2, 1); err != nil {
			log.Errorf(err, "BuyETHHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr

	if buyLen >= 20 {
		log.Infof("BuyETHHandle buyLen more than 20")
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
			log.Errorf(err, "BuyETHHandle Buy error")
			return
		}
		log.Infof("BuyETHHandle success")
		return
	}

}

func SellETHHandle() {
	result, err := pkg.Get(EthUrl, false, "")
	if err != nil {
		log.Errorf(err, "SellETHHandle pkg.Get error")
		return
	}
	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "SellETHHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	sellList := resp.Data.Sell
	sellLen := len(sellList)
	price := int64(resp.Data.Price * 1e6)

	if price >= 30*1e5 || (sellLen > 0 && sellList[sellLen-1].Price*1e6 >= 30*1e5) {
		if err := service.SetRobotType(2, 2); err != nil {
			log.Errorf(err, "SellETHHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr

	if sellLen >= 20 {
		log.Infof("SellETHHandle sellLen more than 20")
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
		err := service.Approve(EthTokenAddr, userAddr, userKey, DexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "SellETHHandle Approve error")
			return
		}
		err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
		if err != nil {
			log.Errorf(err, "sell error")
			return
		}
		log.Infof("SellETHHandle success")
	}

}

func TradeETHHandle() {
	result, err := pkg.Get(EthUrl, false, "")
	if err != nil {
		log.Errorf(err, "TradeETHHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "TradeETHHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	sellList := resp.Data.Sell
	sellLen := len(sellList)

	log.Infof("TradeETHHandle buyLen:%v, sellLen:%v", buyLen, sellLen)

	if buyLen <= 6 || sellLen <= 6 {
		log.Infof("TradeETHHandle buyLen or sellLen less than 6")
		return
	}

	robotType := service.GetRobotType(2)
	if robotType == 0 {
		log.Errorf(nil, "TradeETHHandle obotType is 0")
		return
	}

	log.Infof("TradeETHHandle robotType:%v", robotType)

	currentTime := time.Now().Unix()

	orderType := BUY
	rand := service.RandInt64(1, 101)

	if robotType == 2 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			if err := ethSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 5; i++ {
					ethSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := ethBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 2; i++ {
					ethBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := ethBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 5; i++ {
					ethBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := ethSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 2; i++ {
					ethSell4Supply(sellPrice)
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
			if err := ethBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 5; i++ {
					ethBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := ethSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 2; i++ {
					ethSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := ethSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 5; i++ {
					ethSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := ethBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 2; i++ {
					ethBuy4Supply(buyPrice)
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
		ethBuy(sellList)
	} else {
		ethSell(buyList)
	}
	return
}

func ethBuy(sellList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "ethBuy error")
		return err
	}
	log.Infof("ethBuy success")
	return nil
}

func ethBuy4Five(sellList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	buyPrice := int64(sellList[index-1].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := int64(0)
	for i := 0; i < 5; i++ {
		amount1 += int64(sellList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := service.GetAmount1(buyPrice)
	amount1 += addAmount
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "ethBuy4Five error")
		return err
	}
	log.Infof("ethBuy4Five success")
	return nil
}

func ethSell(buyList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.RandInt64(20, 30) * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EthTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "ethSell Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "ethSell error")
		return err
	}
	log.Infof("ethSell success")
	return nil
}

func ethSell4Five(buyList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	sellPrice := int64(buyList[index-1].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := int64(0)
	for i := 0; i < index; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := service.GetAmount1(sellPrice)
	amount1 += addAmount
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EthTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "ethSell4Five Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "ethSell4Five error")
		return err
	}
	log.Infof("ethSell4Five success")
	return nil
}

func ethTrade4Loop(buyList []*service.PairOrderModel, sellList []*service.PairOrderModel, index int) {
	for i := 0; i < index; i++ {
		ethSell(buyList)
		time.Sleep(500 * time.Millisecond)
		ethBuy(sellList)
		time.Sleep(500 * time.Millisecond)
	}
}

func ethBuy4Supply(buyPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "ethBuy4Supply error")
		return err
	}
	log.Infof("ethBuy4Supply success")
	return nil
}

func ethSell4Supply(sellPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EthTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(sellPrice)
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EthTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "ethSell4Supply Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "ethSell4Supply error")
		return err
	}
	log.Infof("ethSell4Supply success")
	return nil
}
