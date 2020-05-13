package service

import (
	"encoding/json"
	"fmt"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"math/rand"
	"time"
)

const (
	dexContractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	btcTokenAddr    = "TEQEni8FCPrmdTQPUKAu1DCpm3ZYESjFg8"
	trxTokenAddr    = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"
	btcOwner        = "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	btcOwnerKey     = "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07"
	BUY             = 1
	SELL            = 2
)

func BuyBTCHandle() {
	url := fmt.Sprintf("https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=%v", 1)
	result, err := pkg.Get(url, false, "")
	if err != nil {
		log.Errorf(err, "pkg.Get error")
		return
	}

	resp := &ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	buyList := resp.Data.Buy
	buyLen := len(buyList)
	price := int64(resp.Data.Price * 1e6)

	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr

	if buyLen >= 10 {
		log.Infof("buyLen more than 10")
		return
	}

	buyPrice := int64(0)

	if buyLen == 0 {
		tempPrice := price
		if tempPrice <= 700000 {
			tempPrice = 900000
		}
		// 比当前价格少10000
		buyPrice = tempPrice - 10000
	} else if buyLen < 10 {
		tempPrice := int64(buyList[buyLen-1].Price * 1e6)
		if tempPrice <= 700000 {
			tempPrice = 900000
		}
		// 比最近一单价格少1000~1500
		buyPrice = tempPrice - RandInt64(1000, 1500)
	}

	if buyPrice > 0 {
		token2 := trxTokenAddr
		amount1 := RandInt64(20, 30) * 1e6
		amount2 := amount1 * buyPrice / 1e6
		err = Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
		if err != nil {
			log.Errorf(err, "Buy error")
			return
		}
		log.Infof("BuyBTCHandle success")
		return
	}

}

func SellBTCHandle() {
	url := fmt.Sprintf("https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=%v", 1)
	result, err := pkg.Get(url, false, "")
	if err != nil {
		log.Errorf(err, "pkg.Get error")
		return
	}
	resp := &ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "json.Unmarshal error")
		return
	}
	if resp.Code != 0 {
		return
	}

	sellList := resp.Data.Sell
	sellLen := len(sellList)
	price := int64(resp.Data.Price * 1e6)

	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr

	if sellLen >= 10 {
		log.Infof("sellLen more than 10")
		return
	}

	sellPrice := int64(0)

	if sellLen == 0 {
		tempPrice := price
		if tempPrice >= 2200000 {
			tempPrice = 2000000
		}
		// 比当前价格多10000
		sellPrice = tempPrice + 10000
	} else if sellLen < 10 {
		tempPrice := int64(sellList[sellLen-1].Price * 1e6)
		if tempPrice >= 2200000 {
			tempPrice = 2000000
		}
		// 最近一单价格多1000~1500
		sellPrice = tempPrice + RandInt64(1000, 1500)
	}

	if sellPrice > 0 {
		token2 := trxTokenAddr
		amount1 := RandInt64(20, 30) * 1e6
		amount2 := amount1 * sellPrice / 1e6
		err := Approve(btcTokenAddr, userAddr, userKey, dexContractAddr, amount1)
		if err != nil {
			log.Errorf(err, "Approve error")
			return
		}
		err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
		if err != nil {
			log.Errorf(err, "sell error")
			return
		}
		log.Infof("SellBTCHandle success")
	}

}

func TradeBTCHandle() {
	url := fmt.Sprintf("https://bytego123.cn/dex/api/v1/market/pairOrder4Kline/query?pairID=%v", 1)
	result, err := pkg.Get(url, false, "")
	if err != nil {
		log.Errorf(err, "pkg.Get error")
		return
	}

	resp := &ResultResp{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		log.Errorf(err, "json.Unmarshal error")
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

	if buyLen < 5 || sellLen < 5 {
		log.Infof("buyLen or sellLen less than 5")
		return
	}

	currentTime := time.Now().Unix()

	dateTime00 := GetDatetime(currentTime, "1day")

	dateTime12 := dateTime00 + 60*60*(12-8)

	if buyLen >= 20 && buyLen-sellLen >= 5 {
		sell4Five(buyList)
		return
	}

	if sellLen >= 20 && sellLen-buyLen >= 5 {
		buy4Five(sellList)
		return
	}

	orderType := BUY
	rand := RandInt64(1, 101)
	// 12点之前 以卖单为主
	if currentTime <= dateTime12 {
		if currentTime%(30*60) <= 600 {
			sell4Five(buyList)
			return
		}
		if rand <= 15 {
			orderType = BUY
		} else {
			orderType = SELL
		}
	} else if currentTime > dateTime12 {
		if currentTime%(30*60) <= 600 {
			buy4Five(sellList)
			return
		}
		if rand <= 85 {
			orderType = BUY
		} else {
			orderType = SELL
		}
	}

	if orderType == BUY {
		if sellLen >= 18 {
			buy4Five(sellList)
		} else {
			buy(sellList)
		}
	} else {
		if buyLen >= 18 {
			sell4Five(buyList)
		} else {
			sell(buyList)
		}
	}
	return
}

func buy(sellList []*PairOrderModel) error {
	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr
	buyPrice := int64(sellList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * buyPrice / 1e6
	err := Buy(true, userAddr, userKey, token1, token2, amount1, amount2, buyPrice, 0)
	if err != nil {
		log.Errorf(err, "Buy error")
		return err
	}
	log.Infof("TradeBTCHandle buy success")
	return nil
}

func buy4Five(sellList []*PairOrderModel) error {
	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr
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
		log.Errorf(err, "Buy error")
		return err
	}
	log.Infof("TradeBTCHandle buy4Five success")
	return nil
}

func sell(buyList []*PairOrderModel) error {
	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr
	sellPrice := int64(buyList[0].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(btcTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "sell error")
		return err
	}
	log.Infof("TradeBTCHandle sell success")
	return nil
}

func sell4Five(buyList []*PairOrderModel) error {
	userAddr := btcOwner
	userKey := btcOwnerKey
	token1 := btcTokenAddr
	sellPrice := int64(buyList[4].Price * 1e6)
	token2 := trxTokenAddr
	amount1 := RandInt64(20, 30) * 1e6
	for i := 0; i <= 4; i++ {
		amount1 += int64(buyList[i].TotalQuoteAmount * 1e6)
	}
	amount1 += 20 * 1e6
	amount2 := amount1 * sellPrice / 1e6
	err := Approve(btcTokenAddr, userAddr, userKey, dexContractAddr, amount1)
	if err != nil {
		log.Errorf(err, "Approve error")
		return err
	}
	err = Sell(false, userAddr, userKey, token1, token2, amount1, amount2, sellPrice, 0)
	if err != nil {
		log.Errorf(err, "sell error")
		return err
	}
	log.Infof("TradeBTCHandle sell4Five success")
	return nil
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

type ResultResp struct {
	Code int           `json:"code"`
	Data PairOrderResp `json:"data"`
}

type PairOrderResp struct {
	Buy   []*PairOrderModel `json:"buy"`
	Sell  []*PairOrderModel `json:"sell"`
	Price float64           `json:"price"`
}

type PairOrderModel struct {
	Price            float64 `json:"price"`
	TotalQuoteAmount float64 `json:"totalQuoteAmount"`
	TotalBaseAmount  float64 `json:"totalBaseAmount"`
	TotalOrder       int64   `json:"totalOrder"`
}
