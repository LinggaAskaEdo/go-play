package sql

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/xtfly/log4g/api"
)

type SQL interface {
}

type sqlxImpl struct {
	endOnce *sync.Once
	logger  api.Logger
	opt     Options
}

type Options struct {
	DatabaseDriver   string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseHost     string
	DatabasePort     string
}

func Init(logger api.Logger, opt *Options) (*sqlx.DB, error) {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		opt.DatabaseUser,
		opt.DatabasePassword,
		opt.DatabaseHost,
		opt.DatabasePort,
		opt.DatabaseName)

	db, err := sqlx.Connect(opt.DatabaseDriver, databaseURI)
	if err != nil {
		panic(err)
	}

	return db, nil
}
