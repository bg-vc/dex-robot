package transaction

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/service"
	"github.com/vincecfl/go-common/log"
	"gopkg.in/redis.v5"
	"sync"
	"time"
)

const (
	LockRobotTransTrxKey = "dex.robot.trons.trx.task"
)

func TranHandle() {
	for {
		if !lockKey(LockRobotTransTrxKey) {
			log.Infof("TranHandle lock exit, key:%v", LockRobotTransTrxKey)
			time.Sleep(3 * time.Second)
			continue
		}
		err := pkg.RedisCli.Set(LockRobotTransTrxKey, "LockRobotTransValue", 5*time.Minute).Err()
		if err != nil {
			log.Errorf(err, "DetailStatisticTask lock redis set error")
			return
		}
		tranHandleSub()
		time.Sleep(3 * time.Second)
	}

}

func tranHandleSub() {
	defer pkg.RedisCli.Del(LockRobotTransTrxKey)
	url := fmt.Sprintf(`https://apilist.tronscan.org/api/transaction?sort=-timestamp&count=true&limit=10&start=0`)
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	client := fasthttp.Client{}
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	if err := client.DoTimeout(req, resp, 5000*time.Millisecond); err != nil {
		log.Errorf(err, "http.Do error")
		return
	}

	tranResp := &TranResp{}
	if err := json.Unmarshal(resp.Body(), tranResp); err != nil {
		log.Errorf(err, "json.Unmarshal error")
		return
	}

	//log.Infof("tranResp:%v", utils.ToJSONStr(tranResp))

	var wg sync.WaitGroup
	for _, item := range tranResp.DataList {
		wg.Add(1)
		service.TrxHandler(item.Hash, &wg)
	}

	wg.Wait()
	log.Infof("tranHandleSub done")
}

type TranResp struct {
	DataList []*Tran `json:"data"`
}

type Tran struct {
	Block     int64  `json:"block"`
	timestamp int64  `json:"timestamp"`
	Hash      string `json:"hash"`
}

func lockKey(key string) bool {
	value, err := pkg.RedisCli.Get(key).Result()
	if err == redis.Nil {
		return true
	} else if err != nil {
		log.Errorf(err, "lock redis get value error")
		return false
	}
	if value == "" {
		return true
	} else {
		return false
	}
}
