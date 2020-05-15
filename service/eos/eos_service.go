package eos

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"github.com/vincecfl/go-common/log"
	"time"
)

const (
	eosUrl = "https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=3"

	DexContractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	TrxTokenAddr    = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"
	EosTokenAddr    = "TDLqNqjsQkZgyRoyr3t7Fxj45QcXfaMzUu"

	Owner    = "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	OwnerKey = "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07"

	BUY  = 1
	SELL = 2
)

func BuyEOSHandle() {
	result, err := pkg.Get(eosUrl, false, "")
	if err != nil {
		log.Errorf(err, "BuyEOSHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "BuyEOSHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	price := int64(resp.Data.Price * 1e6)

	if price <= 5*1e5 || (buyLen > 0 && buyList[buyLen-1].Price*1e6 <= 5*1e5) {
		if err := service.SetRobotType(3, 1); err != nil {
			log.Errorf(err, "BuyEOSHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr

	if buyLen >= 20 {
		log.Infof("BuyEOSHandle buyLen more than 20")
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
			log.Errorf(err, "BuyEOSHandle Buy error")
			return
		}
		log.Infof("BuyEOSHandle success")
		return
	}

}

func SellEOSHandle() {
	result, err := pkg.Get(eosUrl, false, "")
	if err != nil {
		log.Errorf(err, "SellEOSHandle pkg.Get error")
		return
	}
	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "SellEOSHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	sellList := resp.Data.Sell
	sellLen := len(sellList)
	price := int64(resp.Data.Price * 1e6)

	if price >= 30*1e5 || (sellLen > 0 && sellList[sellLen-1].Price*1e6 >= 30*1e5) {
		if err := service.SetRobotType(3, 2); err != nil {
			log.Errorf(err, "SellEOSHandle SetRobotType error")
			return
		}
	}

	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr

	if sellLen >= 20 {
		log.Infof("SellEOSHandle sellLen more than 20")
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
		err := service.Approve(EosTokenAddr, userAddr, userKey, DexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "SellEOSHandle Approve error")
			return
		}
		err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
		if err != nil {
			log.Errorf(err, "sell error")
			return
		}
		log.Infof("SellEOSHandle success")
	}

}

func TradeEOSHandle() {
	result, err := pkg.Get(eosUrl, false, "")
	if err != nil {
		log.Errorf(err, "TradeEOSHandle pkg.Get error")
		return
	}

	resp := &service.ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "TradeEOSHandle json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	sellList := resp.Data.Sell
	sellLen := len(sellList)

	log.Infof("TradeEOSHandle buyLen:%v, sellLen:%v", buyLen, sellLen)

	if buyLen <= 6 || sellLen <= 6 {
		log.Infof("TradeEOSHandle buyLen or sellLen less than 6")
		return
	}

	robotType := service.GetRobotType(3)
	if robotType == 0 {
		log.Errorf(nil, "TradeEOSHandle obotType is 0")
		return
	}

	log.Infof("TradeEOSHandle robotType:%v", robotType)

	currentTime := time.Now().Unix()

	orderType := BUY
	rand := service.RandInt64(1, 101)

	// robotType为2 卖单为主
	if robotType == 2 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			if err := eosSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 5; i++ {
					eosSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := eosBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 2; i++ {
					eosBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := eosBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 1000 {
				for i := 0; i < 5; i++ {
					eosBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := eosSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 1000 {
				for i := 0; i < 2; i++ {
					eosSell4Supply(sellPrice)
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
			if err := eosBuy4Five(sellList, 5); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[4].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 5; i++ {
					eosBuy4Supply(buyPrice)
					buyPrice = buyPrice - service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 < 2700 {
			if err := eosSell4Five(buyList, 2); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[1].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 2; i++ {
					eosSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 <= 300 && time60 >= 2700 {
			if err := eosSell4Five(buyList, 5); err != nil {
				return
			}
			// 补充卖单
			sellPrice := int64(buyList[4].Price * 1e6)
			if len(sellList) <= 500 {
				for i := 0; i < 5; i++ {
					eosSell4Supply(sellPrice)
					sellPrice = sellPrice + service.RandInt64(2000, 2500)
				}
			}
			return
		} else if time15 > 600 && time15 <= 900 && time60 >= 2700 {
			if err := eosBuy4Five(sellList, 2); err != nil {
				return
			}
			// 补充买单
			buyPrice := int64(sellList[1].Price * 1e6)
			if len(buyList) <= 500 {
				for i := 0; i < 2; i++ {
					eosBuy4Supply(buyPrice)
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
		eosBuy(sellList)
	} else {
		eosSell(buyList)
	}
	return
}

func eosBuy(sellList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "eosBuy error")
		return err
	}
	log.Infof("eosBuy success")
	return nil
}

func eosBuy4Five(sellList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
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
		log.Errorf(err, "eosBuy4Five error")
		return err
	}
	log.Infof("eosBuy4Five success")
	return nil
}

func eosSell(buyList []*service.PairOrderModel) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(sellPrice)
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EosTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "eosSell Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "eosSell error")
		return err
	}
	log.Infof("eosSell success")
	return nil
}

func eosSell4Five(buyList []*service.PairOrderModel, index int) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
	sellPrice := int64(buyList[index-1].Price * 1e6)
	token2 := TrxTokenAddr
	amount1 := int64(0)
	for i := 0; i < index; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := service.GetAmount1(sellPrice)
	amount1 += addAmount
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EosTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "eosSell4Five Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "eosSell4Five error")
		return err
	}
	log.Infof("eosSell4Five success")
	return nil
}

func eosTrade4Loop(buyList []*service.PairOrderModel, sellList []*service.PairOrderModel, index int) {
	for i := 0; i < index; i++ {
		eosSell(buyList)
		time.Sleep(500 * time.Millisecond)
		eosBuy(sellList)
		time.Sleep(500 * time.Millisecond)
	}
}

func eosBuy4Supply(buyPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(buyPrice)
	amount2 := amount1 * buyPrice / 1e6
	err := service.Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "eosBuy4Supply error")
		return err
	}
	log.Infof("eosBuy4Supply success")
	return nil
}

func eosSell4Supply(sellPrice int64) error {
	userAddr := Owner
	userKey := OwnerKey
	token1 := EosTokenAddr
	token2 := TrxTokenAddr
	amount1 := service.GetAmount1(sellPrice)
	amount2 := amount1 * sellPrice / 1e6
	err := service.Approve(EosTokenAddr, userAddr, userKey, DexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "eosSell4Supply Approve error")
		return err
	}
	err = service.Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "eosSell4Supply error")
		return err
	}
	log.Infof("eosSell4Supply success")
	return nil
}
