package service

import (
	"fmt"
	"github.com/vincecfl/dex-robot/pkg"
	"github.com/vincecfl/go-common/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestGetTrx(t *testing.T) {
	pkg.Init("../conf/config.yaml")
	result, err := post("https://www.trongrid.io/api/get-shasta-trx", "TPdBHYrTDiop2fgsmZGDEfNN5SucJADCf4")
	if err != nil {
		log.Errorf(err, "post error")
		return
	}
	log.Infof("result:%v", result)
}

func post(url string, address string) (string, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest error")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("address", address)
	req.Header.Set("search", "")

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	req.Header.Set("Origin", "https://www.trongrid.io")
	req.Header.Set("Referer", "https://www.trongrid.io/shasta")

	response, err := pkg.HttpClient.Do(req)
	if err != nil && response == nil {
		fmt.Println("httpClient.Do error")
		return "", err
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Couldn't parse response body")
			return "", err
		}
		return string(body), nil
	}
}

func postTwo(reqUrl string, address string) (string, error) {
	proxyUrl := "http://27.46.135.203"
	req, _ := http.NewRequest("POST", reqUrl, nil)
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("address", address)
	req.Header.Set("search", "")

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	req.Header.Set("Origin", "https://www.trongrid.io")
	req.Header.Set("Referer", "https://www.trongrid.io/shasta")

	response, err := client.Do(req)
	if err != nil && response == nil {
		fmt.Println("httpClient.Do error")
		return "", err
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Couldn't parse response body")
			return "", err
		}
		return string(body), nil
	}
}
