// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: db-pgx testing

package mcdb

import (
	"fmt"
	"testing"
)
import "github.com/abbeymart/mctest"

func TestDbPgxPool(t *testing.T) {
	// test-data: db-configuration settings
	myDb := DbConfig{
		DbType:   "postgres",
		Host:     "localhost",
		Username: "postgres",
		Password: "ab12testing",
		Port:     5432,
		DbName:   "mcdev",
		Filename: "testdb.db",
		PoolSize: 20,
		Url:      "localhost:5432",
	}

	myDb.Options = DbConnectOptions{}

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the database - pgxPool",
		TestFunc: func() {
			dbc, err := myDb.OpenPgxDbPool()
			defer myDb.ClosePgxDb()
			fmt.Println(dbc)
			fmt.Println("*****************************************")
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	mctest.PostTestResult()
}
