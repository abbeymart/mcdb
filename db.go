// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mcdb - db connection for PostgresSQL, MySQL, SQLite3

package mcdb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type DbConnectionType *sql.DB

type DbSecureType struct {
	SecureAccess bool
	SecureCert   string
	SecureKey    string
	SslMode      string
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
	Url          string
	SecureOption DbSecureType
}

type DbConnectOptions map[string]interface{}

type DbConfig struct {
	DbType       string
	Host         string
	Username     string
	Password     string
	DbName       string
	Filename     string
	Location     string
	Port         uint32
	PoolSize     uint
	Url          string
	SecureOption DbSecureType
	Options      DbConnectOptions
}

var (
	db  *sql.DB
	err error
)

func (dbConfig DbConfig) OpenDb() (*sql.DB, error) {
	sslMode := dbConfig.SecureOption.SslMode
	sslCert := dbConfig.SecureOption.SecureCert
	if sslMode == "" {
		sslMode = "disable"
	}
	switch dbConfig.DbType {
	case "postgres":
		connectionString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=%v sslrootcert=%v", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, sslMode, sslCert)
		//connectionString := fmt.Sprintf("dbname=%v sslmode=disable", dbConfig.Database)
		if os.Getenv("DATABASE_URL") != "" {
			connectionString = os.Getenv("DATABASE_URL")
		}
		db, err = sql.Open(dbConfig.DbType, connectionString)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
			//panic(err)
		}
		return db, nil
	case "mysql", "mariadb":
		connectionString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=%v", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, sslMode)
		if os.Getenv("DATABASE_URL") != "" {
			connectionString = os.Getenv("DATABASE_URL")
		}
		db, err = sql.Open(dbConfig.DbType, connectionString)
		if err != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
		}
		return db, nil
	case "sqlite3":
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
		err = db.Close()
		if err != nil {
			// log error to the console
			fmt.Println(err)
		}
	}
}
