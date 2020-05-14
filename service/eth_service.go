package service

import (
	"encoding/json"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"time"
)

func BuyETHHandle() {
	result, err := pkg.Get(ethUrl, false, "")
	if err != nil {
		log.Errorf(err, "BuyETHHandle pkg.Get error")
		return
	}

	resp := &ResultResp{}
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
		if err := SetRobotType(2, 1); err != nil {
			log.Errorf(err, "BuyETHHandle SetRobotType error")
			return
		}
	}

	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr

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
			log.Errorf(err, "BuyETHHandle Buy error")
			return
		}
		log.Infof("BuyETHHandle success")
		return
	}

}

func SellETHHandle() {
	result, err := pkg.Get(ethUrl, false, "")
	if err != nil {
		log.Errorf(err, "SellETHHandle pkg.Get error")
		return
	}
	resp := &ResultResp{}
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
		if err := SetRobotType(2, 2); err != nil {
			log.Errorf(err, "SellETHHandle SetRobotType error")
			return
		}
	}

	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr

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
		err := Approve(ethTokenAddr, userAddr, userKey, dexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "SellETHHandle Approve error")
			return
		}
		err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
		if err != nil {
			log.Errorf(err, "sell error")
			return
		}
		log.Infof("SellETHHandle success")
	}

}

func TradeETHHandle() {
	result, err := pkg.Get(ethUrl, false, "")
	if err != nil {
		log.Errorf(err, "TradeETHHandle pkg.Get error")
		return
	}

	resp := &ResultResp{}
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

	robotType := GetRobotType(2)
	if robotType == 0 {
		log.Errorf(nil, "TradeETHHandle obotType is 0")
		return
	}

	log.Infof("TradeETHHandle robotType:%v", robotType)

	currentTime := time.Now().Unix()

	orderType := BUY
	rand := RandInt64(1, 101)
	// robotType为2 卖单为主
	if robotType == 2 {
		time60 := currentTime % (60 * 60)
		time15 := currentTime % (15 * 60)
		if time15 <= 300 && time60 < 2700 {
			ethSell4Five(buyList, 5)
			return
		} else if time15 <= 300 && time60 >= 2700 {
			ethBuy4Five(sellList, 2)
			ethTrade4Loop(buyList, sellList, 5)
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
			ethBuy4Five(sellList, 5)
			return
		} else if time15 <= 300 && time60 >= 2700 {
			ethSell4Five(buyList, 2)
			ethTrade4Loop(buyList, sellList, 5)
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

func ethBuy(sellList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * buyPrice / 1e6
	err := Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "ethBuy error")
		return err
	}
	log.Infof("ethBuy success")
	return nil
}

func ethBuy4Five(sellList []*PairOrderModel, index int) error {
	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr
	buyPrice := int64(sellList[index-1].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := int64(0)
	for i := 0; i < 5; i++ {
		amount1 += int64(sellList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := int64(0)
	if buyPrice <= 1*1e6 {
		addAmount = RandInt64(20, 30) * 1e6
	} else if buyPrice > 1*1e6 && buyPrice <= 2*1e6 {
		addAmount = RandInt64(10, 15) * 1e6
	} else if buyPrice > 2*1e6 {
		addAmount = RandInt64(5, 10) * 1e6
	}
	amount1 += addAmount
	amount2 := amount1 * buyPrice / 1e6
	err := Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "ethBuy4Five error")
		return err
	}
	log.Infof("ethBuy4Five success")
	return nil
}

func ethSell(buyList []*PairOrderModel) error {
	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(ethTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "ethSell Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "ethSell error")
		return err
	}
	log.Infof("ethSell success")
	return nil
}

func ethSell4Five(buyList []*PairOrderModel, index int) error {
	userAddr := owner
	userKey := ownerKey
	token1 := ethTokenAddr
	sellPrice := int64(buyList[index-1].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	for i := 0; i < index; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	addAmount := int64(0)
	if sellPrice <= 1*1e6 {
		addAmount = RandInt64(20, 30) * 1e6
	} else if sellPrice > 1*1e6 && sellPrice <= 2*1e6 {
		addAmount = RandInt64(10, 15) * 1e6
	} else if sellPrice > 2*1e6 {
		addAmount = RandInt64(5, 10) * 1e6
	}
	amount1 += addAmount
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(ethTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "ethSell4Five Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "ethSell4Five error")
		return err
	}
	log.Infof("ethSell4Five success")
	return nil
}

func ethTrade4Loop(buyList []*PairOrderModel, sellList []*PairOrderModel, index int) {
	for i := 0; i < index; i++ {
		btcSell(buyList)
		time.Sleep(500 * time.Millisecond)
		btcBuy(sellList)
		time.Sleep(500 * time.Millisecond)
	}
}
