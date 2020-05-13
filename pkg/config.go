package pkg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/vincecfl/go-common/log"
	"gopkg.in/redis.v5"
	"net/http"
)

var (
	Cfg        *Config
	DBWrite    *sqlx.DB
	DBRead     *sqlx.DB
	RedisCli   *redis.Client
	HttpClient *http.Client
)

type Config struct {
	Logger     *LoggerConfig
	HttpClient *HttpClientConfig
}

func Init(cfgName string) {
	setConfig(cfgName)
	Cfg = loadConfig()
	initConfig(Cfg)
	watchConfig()
}

func setConfig(cfgName string) {
	if cfgName != "" {
		viper.SetConfigFile(cfgName)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("initConfig error")
	}
}

func loadConfig() *Config {
	cfg := &Config{
		Logger:     LoadLoggerConfig(viper.Sub("logger")),
		HttpClient: LoadHttpClientConfig(),
	}
	return cfg
}

func initConfig(cfg *Config) {
	cfg.Logger.InitLogger()
	HttpClient = cfg.HttpClient.InitHttpClient()
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}
