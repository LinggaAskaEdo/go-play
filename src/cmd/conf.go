package main

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	// Log Config
	LogConfigPath string `mapstructure:"LOG_CONFIG_PATH"`
	LogConfigName string `mapstructure:"LOG_CONFIG_NAME"`

	// Database Config
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseUrl      string `mapstructure:"DATABASE_URL"`
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
