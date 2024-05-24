package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/linggaaskaedo/go-play/src/business/domain"
	"github.com/linggaaskaedo/go-play/src/business/usecase"
	"github.com/linggaaskaedo/go-play/src/business/usecase/rss"
	schedulerhandler "github.com/linggaaskaedo/go-play/src/handler/scheduler"

	// grace "github.com/linggaaskaedo/go-play/stdlib/grace"
	log "github.com/linggaaskaedo/go-play/stdlib/logger"
	libsql "github.com/linggaaskaedo/go-play/stdlib/sql"
)

var (
	// Resource
	logger     log.Logger
	sqlClient0 libsql.SQL
	// app        grace.App

	// Business Layer
	dom *domain.Domain
	uc  *usecase.Usecase

	// Handlers
	scheduler *schedulerhandler.Scheduler
)

// func init() {
func main() {
	// Initialize environment configurations
	InitEnvConfigs()

	// Add Sleep with Jitter to drag the the initialization time among instances
	SleepWithJitter()

	// Set up logging configuration using environment variables
	logConfig := log.LogConfig{
		LogConfigPath: EnvConfigs.LogConfigPath,
		LogConfigName: EnvConfigs.LogConfigName,
	}

	logger, err := log.Init(&logConfig)
	if err != nil {
		panic(err)
	}

	// Set up database configuration using environment variables
	databaseConfig := libsql.DBConfig{
		DatabaseUser:     EnvConfigs.DatabaseUser,
		DatabasePassword: EnvConfigs.DatabasePassword,
		DatabaseName:     EnvConfigs.DatabaseName,
		DatabaseUrl:      EnvConfigs.DatabaseUrl,
		DatabasePort:     EnvConfigs.DatabasePort,
	}

	// Establish a connection to the MySQL database
	sqlClient0, err := libsql.Init(
		logger,
		&databaseConfig,
	)
	if err != nil {
		panic(err)
	}

	// Construct domain
	dom = domain.Init(
		logger,
		sqlClient0,
	)

	// Construct usecase
	uc = usecase.Init(
		logger,
		usecase.Options{
			RSS: rss.Options{
				RSSUrl: EnvConfigs.RSSUrl,
			},
		},
		dom,
	)

	// Scheduler Handler Initialization
	schedulerhandler.Init(
		logger, schedulerhandler.Options{
			RSS: schedulerhandler.RSSOptions{
				Enabled: EnvConfigs.RSSSchedulerEnable,
				Period:  EnvConfigs.RSSSchedulerCron,
			},
		},
		uc,
	)

	// Keep running application until terminate
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Println("Application Terminated !!!")
}
