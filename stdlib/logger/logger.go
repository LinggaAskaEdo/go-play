package logger

import (
	"fmt"

	log "github.com/xtfly/log4g"
	"github.com/xtfly/log4g/api"
)

type Logger interface {
}

type Options struct {
	LogConfigPath string
	LogConfigName string
}

func Init(opt *Options) (api.Logger, api.Manager, error) {
	err := log.GetManager().LoadConfigFile(opt.LogConfigPath)
	if err != nil {
		fmt.Printf("Error loading configuration log: %v\nConfig: %+v\n", err, opt)
		panic(err)
	}

	dlog := log.GetLogger(opt.LogConfigName)
	mlog := log.GetManager()

	return dlog, mlog, nil
}
