// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: mcdb - db connection for PostgresSQL (pgx)

package mcdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var (
	dbpgxP  *pgxpool.Pool
	errpgxP error
)

type DbPool struct {
	DbConn *pgxpool.Pool
}

func (dbConfig DbConfig) OpenPgxDbPool() (*DbPool, error) {
	switch dbConfig.DbType {
	case "postgres":
		connectionString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Port, dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName)
		if os.Getenv("DATABASE_URL") != "" {
			connectionString = os.Getenv("DATABASE_URL")
		}
		// parseConfig
		config, err := pgxpool.ParseConfig(connectionString)
		if err != nil {
			errMsg := fmt.Sprintf("Parsing Connection Configuration Error: %v", err)
			return nil, errors.New(errMsg)
		}
		// perform db-configuration tasks
		dbpgxP, errpgxP = pgxpool.ConnectConfig(context.Background(), config)
		if errpgxP != nil {
			errMsg := fmt.Sprintf("Database Connection Error: %v", err)
			return nil, errors.New(errMsg)
			//panic(err)
		}
		return &DbPool{DbConn: dbpgxP}, nil
	default:
		return nil, errors.New("unknown db-type('postgres')")
	}
}

func (dbConfig DbConfig) ClosePgxDbPool() {
	if dbpgxP != nil {
		dbpgxP.Close()
	}
}
