package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xtfly/log4g/api"
)

type SQL interface{}

type DBConfig struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseUrl      string
	DatabasePort     string
}

// ConnectMySQL- Connect to MySQL database
func Init(logger api.Logger, config *DBConfig) (*sql.DB, error) {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseUrl,
		config.DatabasePort,
		config.DatabaseName)

	db, err := sql.Open("mysql", databaseURI)
	if err != nil {
		logger.Errorf("Error connecting to database: %v\nConfig: %+v\n", err, config)
		return nil, err
	}

	dbErr := db.Ping()
	if dbErr != nil {
		logger.Errorf("Error ping to database: %v\n", dbErr.Error())
		return nil, dbErr
	}

	logger.Info("Successfully connected to database: ", config.DatabaseName)

	return db, nil
}
