package service

import (
	"encoding/json"
	"fmt"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"math/rand"
)

const (
	dexContractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	btcTokenAddr    = "TEQEni8FCPrmdTQPUKAu1DCpm3ZYESjFg8"
	trxTokenAddr    = "T9yD14Nj9j7xAB4dbGeiX9h8unkKHxuWwb"
	btcOwner        = "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4"
	btcOwnerKey     = "514bfc62a1f84b69a46ba6478f991eacb136ef1a2f63a16a66e7f42c14c1de07"
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
		// 比最近一单价格少100
		buyPrice = tempPrice - RandInt64(100, 150)
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
		// 最近一单价格多100
		sellPrice = tempPrice + RandInt64(100, 150)
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

	if buyLen < 5 || sellLen < 5 {
		log.Infof("buyLen or sellLen less than 5")
		return
	}

	if buyLen-sellLen >= 5 {
		if err := sell(buyList); err != nil {
			log.Errorf(err, "sell error")
		}
		return
	}

	if sellLen-buyLen >= 5 {
		if err := buy(sellList); err != nil {
			log.Errorf(err, "buy error")
		}
		return
	}

	rand := RandInt64(1, 3)
	if rand == 1 {
		if err := buy(sellList); err != nil {
			log.Errorf(err, "buy error")
		}
	} else {
		if err := sell(buyList); err != nil {
			log.Errorf(err, "sell error")
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
