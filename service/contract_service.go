package service

import (
	"encoding/json"
	"errors"
	"github.com/tronprotocol/grpc-gateway/api"
	"github.com/tronprotocol/grpc-gateway/core"
	"github.com/vincecfl/go-common/log"
	"github.com/vincecfl/go-common/tron/common"
	"github.com/vincecfl/go-common/tron/tools"
	"github.com/vincecfl/go-common/tron/utils"
)

var (
	contractAddr = "TJ86JLUrMEXYQPNXx1tyD1SzxEgPECFpmj"
	feeLimit     = int64(10000000)
	test         = 1

	buyMethod    *tools.Method
	sellMethod   *tools.Method
	cancelMethod *tools.Method
)

func InitSmart() error {
	if test == 1 {
		utils.TestNet = true
		utils.NetName = utils.NetShasta
	} else {
	}
	smartContract, err := common.GetWalletClient().GetContract(contractAddr)
	abi, err := json.Marshal(smartContract.Abi)
	if nil != err {
		return err
	}
	trade, err := tools.GetABI(string(abi))
	if nil != err {
		log.Errorf(err, "tools.GetABI error")
		return err
	}

	buyMethod = trade.GetMethod("buy")
	if nil == buyMethod {
		return errors.New("buyMethod nil")
	}
	sellMethod = trade.GetMethod("sell")
	if nil == sellMethod {
		return errors.New("sellMethod nil")
	}
	cancelMethod = trade.GetMethod("cancel")
	if nil == cancelMethod {
		return errors.New("cancelMethod nil")
	}

	return nil
}

func Buy(isTrxBase bool, userAddr, userKey string, quoteToken, baseToken string, quoteAmount, baseAmount, price, channelID int64) error {
	record := &CallRecord{}
	record.Data, record.Err = buyMethod.Pack(tools.GenAbiAddress(quoteToken), tools.GenAbiInt(quoteAmount),
		tools.GenAbiAddress(baseToken), tools.GenAbiInt(baseAmount), tools.GenAbiInt(price), tools.GenAbiInt(channelID))
	if nil != record.Err {
		log.Errorf(record.Err, "pack error")
		return record.Err
	}
	callValue := int64(0)
	if isTrxBase {
		callValue = baseAmount
	}
	ctxType, ctx, err := tools.GenTriggerSmartContract(userAddr, contractAddr, callValue, record.Data)
	if err != nil {
		log.Errorf(record.Err, "GenTriggerSmartContract error")
		return err
	}

	trxHash, result, err := broadcastCtxWithFeeLimit(ctxType, ctx, userKey, feeLimit)
	if nil != err || result == nil {
		log.Errorf(err, "broadcastCtxWithFeeLimit error")
		return errors.New("broadcast error")
	}

	log.Infof("hash:[%v]-->result.code:%v-->%s", trxHash, result.Code.String(), result.Message)
	return nil
}

func Sell(isTrxQuote bool, userAddr, userKey string, quoteToken, baseToken string, quoteAmount, baseAmount, price, channelID int64) error {
	record := &CallRecord{}
	record.Data, record.Err = sellMethod.Pack(tools.GenAbiAddress(quoteToken), tools.GenAbiInt(quoteAmount),
		tools.GenAbiAddress(baseToken), tools.GenAbiInt(baseAmount), tools.GenAbiInt(price), tools.GenAbiInt(channelID))
	if nil != record.Err {
		log.Errorf(record.Err, "pack error")
		return record.Err
	}
	callValue := int64(0)
	if isTrxQuote {
		callValue = quoteAmount
	}
	ctxType, ctx, err := tools.GenTriggerSmartContract(userAddr, contractAddr, callValue, record.Data)
	if err != nil {
		log.Errorf(record.Err, "GenTriggerSmartContract error")
		return err
	}
	trxHash, result, err := broadcastCtxWithFeeLimit(ctxType, ctx, userKey, feeLimit)
	if nil != err || result == nil {
		log.Errorf(err, "broadcastCtxWithFeeLimit error")
		return errors.New("broadcast error")
	}

	log.Infof("hash:[%v]-->result.code:%v-->%s", trxHash, result.Code.String(), result.Message)
	return nil
}

func Cancel(userAddr, userKey string, orderID int64) error {
	record := &CallRecord{}
	record.Data, record.Err = cancelMethod.Pack(tools.GenAbiInt(orderID))
	if nil != record.Err {
		log.Errorf(record.Err, "pack error")
		return record.Err
	}
	ctxType, ctx, err := tools.GenTriggerSmartContract(userAddr, contractAddr, 0, record.Data)
	if err != nil {
		log.Errorf(record.Err, "GenTriggerSmartContract error")
		return err
	}

	trxHash, result, err := broadcastCtxWithFeeLimit(ctxType, ctx, userKey, feeLimit)
	if nil != err || result == nil {
		log.Errorf(err, "broadcastCtxWithFeeLimit error")
		return errors.New("broadcast error")
	}

	log.Infof("hash:[%v]-->result.code:%v-->%s", trxHash, result.Code.String(), result.Message)
	return nil
}

func Approve(contractAddr, OwnerAddr, OwnerKey, spender string, value int64) error {
	smartContract, err := common.GetWalletClient().GetSmartContract(contractAddr)
	abi, err := json.Marshal(smartContract.Abi)
	if nil != err {
		log.Errorf(err, "json.Marshal error")
		return err
	}
	token, err := tools.GetABI(string(abi))
	if nil != err {
		log.Errorf(err, "tools.GetABI")
		return err
	}
	approveMethod := token.GetMethod("approve")
	if nil == approveMethod {
		return errors.New("approveMethod is nil")
	}
	data, err := approveMethod.Pack(tools.GenAbiAddress(spender), tools.GenAbiInt(value))
	if nil != err {
		log.Errorf(err, "approveMethod.Pack")
		return err
	}
	ctxType, ctx, err := tools.GenTriggerSmartContract(OwnerAddr, contractAddr, 0, data)
	if err != nil {
		log.Errorf(err, "GenTriggerSmartContract error")
		return err
	}

	trxHash, result, err := broadcastCtxWithFeeLimit(ctxType, ctx, OwnerKey, feeLimit)
	if nil != err || result == nil {
		log.Errorf(err, "broadcastCtxWithFeeLimit error")
		return errors.New("broadcast error")
	}

	log.Infof("hash:%v,result.code:%v-->%s", trxHash, result.Code.String(), result.Message)

	return nil
}

func broadcastCtxWithFeeLimit(ctxType core.Transaction_Contract_ContractType, ctx interface{}, privateKey string, feeLimit int64) (string, *api.Return, error) {
	var hash string
	var result *api.Return
	var err error
	tryCnt := 30
	for tryCnt > 0 {
		tryCnt--
		hash, result, err = tools.BroadcastCtxWithFeeLimit(ctxType, ctx, nil, privateKey, 0, feeLimit)
		if err != nil || result == nil {
			log.Errorf(err, "BroadcastCtxWithFeeLimit error")
			continue
		}
		return hash, result, err
	}
	return hash, result, err
}

type CallRecord struct {
	Owner     string // base58
	Contract  string // base58
	Method    string
	CallValue int64
	Data      []byte
	TrxHash   string // hex
	Err       error
	Return    *api.Return
}
