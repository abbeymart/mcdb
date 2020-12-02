// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcdb

import (
	"database/sql"
	"errors"
	"fmt"
)

type DbConnectionType *sql.DB

type DbSecureType struct {
	secureAccess bool
	secureCert   string
	secureKey    string
}

type DbConfigType struct {
	host         string
	username     string
	password     string
	database     string
	filename     string
	location     string
	port         uint32
	dbType       string
	poolSize     uint
	secureOption DbSecureType
	uri          string
}

type DbConnectOptions map[string]interface{}

type DbConfig struct {
	dbType   string
	dbConfig DbConfigType
	options  DbConnectOptions
}

var (
	db  DbConnectionType
	err error
)

func (dbInfo DbConfig) OpenDb() (DbConnectionType, error) {
	switch dbInfo.dbType {
	case "postgres":
		dataSourceName := fmt.Sprintf("dbname=%v sslmode=disable", dbInfo.dbConfig.database)
		db, err = sql.Open(dbInfo.dbType, dataSourceName)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
			//panic(err)
		}
		return db, nil
	case "mysql":
		dataSourceName := fmt.Sprintf("dbname=%v sslmode=disable", dbInfo.dbConfig.database)
		db, err = sql.Open(dbInfo.dbType, dataSourceName)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
		}
		return db, nil
	case "sqlite":
		dataSourceName := fmt.Sprintf("dbname=%v sslmode=disable", dbInfo.dbConfig.database)
		db, err = sql.Open(dbInfo.dbType, dataSourceName)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
		}
		return db, nil
	default:
		return nil, errors.New("unknown db-type('postgres')")
	}
}
