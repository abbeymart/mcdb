// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcdb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type DbConnectionType *sql.DB

type DbSecureType struct {
	SecureAccess bool
	SecureCert   string
	SecureKey    string
}

type DbConfigType struct {
	Host         string
	Username     string
	Password     string
	DbName       string
	Filename     string
	Location     string
	Port         uint32
	DbType       string
	PoolSize     uint
	SecureOption DbSecureType
	Uri          string
}

type DbConnectOptions map[string]interface{}

type DbConfig struct {
	DbType string
	DbConfigType
	Options DbConnectOptions
}

var (
	db  *sql.DB
	err error
)

func (dbConfig DbConfig) OpenDb() (DbConnectionType, error) {
	switch dbConfig.DbType {
	case "postgres":
		dataSourceName := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName)
		//dataSourceName := fmt.Sprintf("dbname=%v sslmode=disable", dbConfig.Database)
		db, err = sql.Open(dbConfig.DbType, dataSourceName)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
			//panic(err)
		}
		return db, nil
	case "mysql", "mariadb":
		dataSourceName := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName)
		db, err = sql.Open(dbConfig.DbType, dataSourceName)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
		}
		return db, nil
	case "sqlite3":
		//dataSourceName := fmt.Sprintf("dbname=%v sslmode=disable", dbConfig.Database)
		db, err = sql.Open(dbConfig.DbType, dbConfig.Filename)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
		}
		return db, nil
	default:
		return nil, errors.New("unknown db-type('postgres')")
	}
}

func (dbConfig DbConfig) CloseDb() {
	if db != nil {
		_ = db.Close()
	}
}
