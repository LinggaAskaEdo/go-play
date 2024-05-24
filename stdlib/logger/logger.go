package logger

import (
	"fmt"

	log "github.com/xtfly/log4g"
	"github.com/xtfly/log4g/api"
)

type Logger interface{}

type LogConfig struct {
	LogConfigPath string
	LogConfigName string
}

// Initialize log configuration
func Init(config *LogConfig) (api.Logger, error) {
	err := log.GetManager().LoadConfigFile(config.LogConfigPath)
	if err != nil {
		fmt.Printf("Error loading configuration log: %v\nConfig: %+v\n", err, config)
		return nil, err
	}

	dlog := log.GetLogger(config.LogConfigName)

	return dlog, nil
}
