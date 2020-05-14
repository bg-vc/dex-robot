package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/vincecfl/go-common/log"
	"gopkg.in/redis.v5"
	"time"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

func LoadRedisConfig(viper *viper.Viper) *RedisConfig {
	cfg := &RedisConfig{
		Addr:     viper.GetString("host"),
		Password: viper.GetString("pass"),
		DB:       viper.GetInt("index"),
		PoolSize: viper.GetInt("pool_size"),
	}
	fmt.Printf("LoadRedisConfig:%#v\n", cfg)
	return cfg
}

func (cfg *RedisConfig) InitRedis() *redis.Client {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	if redisCli == nil {
		panic("InitRedis RedisCli is nil")
	}
	_, err := redisCli.Ping().Result()
	if err != nil {
		panic("InitRedis error")
	}
	return redisCli
}

var MaxErrCnt = 5

func SetRedisVal(key string, val interface{}, expiration time.Duration) error {
	buf, err := json.Marshal(val)
	if err != nil {
		log.Errorf(err, "json.Marshal error")
		return err
	}

	cmd := RedisCli.Set(key, string(buf), expiration)
	if cmd.Err() != nil {
		log.Errorf(cmd.Err(), "RedisCli.Set error")
		return cmd.Err()
	}
	return nil
}

func GetRedisVal(key string) string {
	tryCnt := 0
	for {
		value, err := RedisCli.Get(key).Result()
		if err == redis.Nil {
			return ""
		} else if err != nil {
			tryCnt++
			if tryCnt > MaxErrCnt {
				log.Errorf(err, "RedisCli.Get error")
				return ""
			}
			continue
		}
		return value
	}
}

func RedisExists(key string) bool {
	errCnt := 0
	for {
		ret := RedisCli.Exists(key)
		if ret == nil || ret.Err() != nil {
			errCnt++
		} else {
			return ret.Val()
		}

		if errCnt > MaxErrCnt {
			log.Error("redisExists error", nil)
			break
		}
	}

	return false
}
