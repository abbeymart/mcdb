// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mcdb - db connection for PostgresSQL (pgx)

package mcdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var (
	dbpgx  *pgx.Conn
	errpgx error
)

type Db struct {
	DbConn *pgx.Conn
}

func (dbConfig DbConfig) OpenPgxDb() (*Db, error) {
	sslMode := dbConfig.SecureOption.SslMode
	sslCert := dbConfig.SecureOption.SecureCert
	if sslMode == "" {
		sslMode = "prefer"
	}
	switch dbConfig.DbType {
	case "postgres":
		connectionString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=%v sslrootcert=%v", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, sslMode, sslCert)
		if os.Getenv("DATABASE_URL") != "" {
			connectionString = os.Getenv("DATABASE_URL")
		}
		// parseConfig
		config, err := pgx.ParseConfig(connectionString)
		if err != nil {
			errMsg := fmt.Sprintf("Parsing Connection Configuration Error: %v", err)
			return nil, errors.New(errMsg)
		}
		// perform db-configuration tasks
		dbpgx, errpgx = pgx.ConnectConfig(context.Background(), config)
		if errpgx != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
			//panic(err)
		}
		return &Db{DbConn: dbpgx}, nil
	default:
		return nil, errors.New("unknown db-type('postgres')")
	}
}

func (dbConfig DbConfig) ClosePgxDb() {
	if dbpgx != nil {
		err = dbpgx.Close(context.Background())
		if err != nil {
			// log error to the console
			fmt.Println(err)
		}
	}
}
