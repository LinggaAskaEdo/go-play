package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

const (
	DefaultMaxJitter = 2000
	DefaultMinJitter = 100
)

type envConfigs struct {
	// Log Config
	LogConfigPath string `mapstructure:"LOG_CONFIG_PATH"`
	LogConfigName string `mapstructure:"LOG_CONFIG_NAME"`

	// Database Config
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseHost      string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`

	// RSS Config
	RSSUrl             string `mapstructure:"RSS_URL"`
	RSSSchedulerEnable bool   `mapstructure:"RSS_SCHEDULER_ENABLE"`
	RSSSchedulerCron   string `mapstructure:"RSS_SCHEDULER_CRON"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

func SleepWithJitter() {
	min := DefaultMinJitter
	max := DefaultMaxJitter

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	rnd := rng.Intn(max-min) + min
	time.Sleep(time.Duration(rnd) * time.Millisecond)
}
