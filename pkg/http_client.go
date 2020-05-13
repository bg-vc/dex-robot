package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

type HttpClientConfig struct {
	TimeOut             time.Duration
	KeepAlive           time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

func LoadHttpClientConfig() *HttpClientConfig {
	cfg := &HttpClientConfig{
		TimeOut:             30,
		KeepAlive:           30,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     30,
	}
	fmt.Printf("LoadHttpClientConfig:%#v\n", cfg)
	return cfg
}

func (cfg *HttpClientConfig) InitHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.TimeOut * time.Second,
				KeepAlive: cfg.KeepAlive * time.Second,
			}).DialContext,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			IdleConnTimeout:     cfg.IdleConnTimeout * time.Second,
		},
		Timeout: cfg.TimeOut * time.Second,
	}
	return client
}

func Get(url string, withToken bool, token string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest error")
	}
	req.Header.Set("Content-Type", "application/json")
	if withToken {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	response, err := HttpClient.Do(req)
	if err != nil && response == nil {
		fmt.Println("HttpClient.Do error")
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

func Post(url string, data interface{}) (string, error) {
	value, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.Marshal error")
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(value)))
	if err != nil {
		fmt.Println("http.NewRequest error")
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := HttpClient.Do(req)
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
