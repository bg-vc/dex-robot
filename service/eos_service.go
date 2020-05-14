package service

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"time"
)

func BuyEOSHandle() {
	result, err := pkg.Get(eosUrl, false, "")
	if err != nil {
		log.Errorf(err, "BuyEOSHandle pkg.Get error")
		return
	}

	resp := &ResultResp{}
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
		if err := SetRobotType(3, 1); err != nil {
			log.Errorf(err, "BuyEOSHandle SetRobotType error")
			return
		}
	}

	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr

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
		// 比最近一单价格少1000~1500
		buyPrice = tempPrice - RandInt64(1000, 1500)
	}

	if buyPrice > 0 {
		token2 := trxTokenAddr
		amount1 := int64(0)
		if buyPrice <= 1*1e6 {
			amount1 = RandInt64(20, 30) * 1e6
		} else if buyPrice > 1*1e6 && buyPrice <= 2*1e6 {
			amount1 = RandInt64(10, 15) * 1e6
		} else if buyPrice > 2*1e6 {
			amount1 = RandInt64(5, 10) * 1e6
		}
		amount2 := amount1 * buyPrice / 1e6
		err = Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
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
	resp := &ResultResp{}
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
		if err := SetRobotType(3, 2); err != nil {
			log.Errorf(err, "SellEOSHandle SetRobotType error")
			return
		}
	}

	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr

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
		// 最近一单价格多1000~1500
		sellPrice = tempPrice + RandInt64(1000, 1500)
	}

	if sellPrice > 0 {
		token2 := trxTokenAddr
		amount1 := int64(0)
		if sellPrice <= 1*1e6 {
			amount1 = RandInt64(20, 30) * 1e6
		} else if sellPrice > 1*1e6 && sellPrice <= 2*1e6 {
			amount1 = RandInt64(10, 15) * 1e6
		} else if sellPrice > 2*1e6 {
			amount1 = RandInt64(5, 10) * 1e6
		}
		amount2 := amount1 * sellPrice / 1e6
		err := Approve(eosTokenAddr, userAddr, userKey, dexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "SellEOSHandle Approve error")
			return
		}
		err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
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

	resp := &ResultResp{}
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

	robotType := GetRobotType(3)
	if robotType == 0 {
		log.Errorf(nil, "TradeEOSHandle obotType is 0")
		return
	}

	log.Infof("TradeEOSHandle robotType:%v", robotType)

	currentTime := time.Now().Unix()

	orderType := BUY
	rand := RandInt64(1, 101)
	// robotType为2 卖单为主
	if robotType == 2 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			eosSell4Five(buyList)
			return
		} else if time15 <= 300 && time60 >= 2700 {
			eosBuy4Five(sellList)
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
			eosBuy4Five(sellList)
			return
		} else if time15 <= 300 && time60 >= 2700 {
			eosSell4Five(buyList)
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

func eosBuy(sellList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * buyPrice / 1e6
	err := Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "eosBuy error")
		return err
	}
	log.Infof("eosBuy success")
	return nil
}

func eosBuy4Five(sellList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr
	buyPrice := int64(sellList[4].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := int64(0)
	for i := 0; i <= 4; i++ {
		amount1 += int64(sellList[i].TotalQuoteAmount * 1e6)
	}
	amount1 += 20 * 1e6
	amount2 := amount1 * buyPrice / 1e6
	err := Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "eosBuy4Five error")
		return err
	}
	log.Infof("eosBuy4Five success")
	return nil
}

func eosSell(buyList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(eosTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "eosSell Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "eosSell error")
		return err
	}
	log.Infof("eosSell success")
	return nil
}

func eosSell4Five(buyList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := eosTokenAddr
	sellPrice := int64(buyList[4].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	for i := 0; i <= 4; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	amount1 += 20 * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(eosTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "eosSell4Five Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "eosSell4Five error")
		return err
	}
	log.Infof("eosSell4Five success")
	return nil
}
