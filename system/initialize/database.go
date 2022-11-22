package initialize

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/config"
	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
)

// LoadAllDatabaseConnection load All Database Connection
func LoadAllDatabaseConnection(h *handler.Handler) error {
	cfg, err := h.GetConfig()
	if err != nil {
		return err
	}

	if cfg != nil {
		dbs := cfg.Databases
		fmt.Printf("dbs: %v\n", dbs)
		e := reflect.ValueOf(&dbs).Elem()
		for i := 0; i < e.NumField(); i++ {
			// varName := e.Type().Field(i).Name
			// varType := e.Type().Field(i).Type
			dbConfig := e.Field(i).Interface()
			// fmt.Printf("%v %v %v\n", varName, varType, dbConfig)
			if dbConfig != nil {
				err := LoadDatabaseConnection(dbConfig.(config.Database), h)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// LoadDatabaseConnection load Database Connection using GORM
func LoadDatabaseConnection(dbConfig config.Database, h *handler.Handler) error {
	fmt.Printf("dbConfig: %v\n", dbConfig)
	var retries = 0
	if h != nil {
		connString := dbConfig.Username + ":" + dbConfig.Password + "@(" + dbConfig.HostName + ")/" + dbConfig.DBName + "?" + dbConfig.Config
		if dbConfig.Driver == "sqlite3" {
			connString = dbConfig.HostName
		}

		dbCon, err := openDBConnection(dbConfig.Driver, connString)

		for err != nil {
			// if logger != nil {
			// 	logger(err, fmt.Sprintf("Failed to connect to database (%d)", retries))
			// }
			fmt.Sprintf("Failed to connect to database (%d)", retries)
			time.Sleep(5 * time.Second)
			dbCon, err = openDBConnection(dbConfig.Driver, connString)
			continue

		}

		dbClient, err := dbCon.DB()
		if err != nil {
			return err
		}
		dbClient.SetMaxIdleConns(dbConfig.MaxIdleConns)
		dbClient.SetMaxOpenConns(dbConfig.MaxOpenConns)
		if dbConfig.LogMode {
			dbCon.Logger = dbCon.Logger.LogMode(logger.Info)
		}

		err = dbClient.Ping()
		if err != nil {
			// h.GetLogger().Errorf("%s Ping error: [%s]", dbConfig.Driver, err.Error())
			return err
		}

		h.SetGormDB(dbConfig.ConnectionName, dbCon)
	}

	return nil
}

func openDBConnection(driverName, dataSourceName string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	if driverName == "postgres" {
		db, err = gorm.Open(postgres.Open(dataSourceName+" dbname=postgres"), &gorm.Config{})
	} else if driverName == "mysql" {
		db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	} else if driverName == "sqlite3" {
		db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config{})
	} else if driverName == "sqlserver" {
		db, err = gorm.Open(sqlserver.Open(dataSourceName), &gorm.Config{})
	} else {
		return nil, errors.New("database dialect is not supported")
	}
	if err != nil {
		return nil, err
	}
	return db, err
}

// CloseDBConnections close DB Connections
func CloseDBConnections(h *handler.Handler) {
	gorms := h.GetGormDBs()
	for key, db := range gorms {
		dbCon, _ := db.DB()
		fmt.Printf("Closing DB Connection `%s`\n", key)
		if err := dbCon.Close(); err != nil {
			fmt.Printf("Error while closing DB Connection `%s`: %s\n", key, err.Error())
		}
	}
}
