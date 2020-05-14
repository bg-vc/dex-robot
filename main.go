package main

import (
	"crypto/tls"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/dex-robot/router"
	"github.com/vincecfl/dex-robot/service"
	"github.com/vincecfl/dex-robot/task"
	"github.com/vincecfl/go-common/log"
	"net/http"
	"time"
)

var (
	cfg = pflag.StringP("config", "c", "", "server config file path")
)

func main() {
	pflag.Parse()

	pkg.Init(*cfg)

	gin.SetMode(viper.GetString("run_mode"))

	g := gin.New()

	router.Load(
		g,
	)

	service.InitSmart()

	go task.BuyBTCHandle()
	go task.SellBTCHandle()
	go task.TradeBTCHandle()

	go task.BuyETHHandle()
	go task.SellETHHandle()
	go task.TradeETHHandle()

	//go task.BuyEOSHandle()
	//go task.SellEOSHandle()
	//go task.TradeEOSHandle()

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("the router has no response, or it might took too long to start up", err)
		}
		log.Info("the router has been deployed successfully")
	}()

	log.Infof("start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get(viper.GetString("url") + "/dex/api/check/health")

		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("connect to the router error")
}
