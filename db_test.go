// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: db testing

package mcdb

import (
	"fmt"
	"testing"
)
import "github.com/abbeymart/mctest"

func TestDb(t *testing.T) {
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
		SecureOption: DbSecureType{
			SslMode: "verify-full",
		},
	}

	myDb.Options = DbConnectOptions{}

	sqliteDb := DbConfig{
		DbType: "sqlite3",
	}
	sqliteDb.Filename = "testdb.db"

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the PostgresDB",
		TestFunc: func() {
			dbc, err := myDb.OpenDb()
			//fmt.Printf("pg-dbc: %v\n", dbc)
			//fmt.Printf("pg-dbc-error: %v\n", err)
			defer myDb.CloseDb()
			fmt.Println(dbc)
			fmt.Println("*****************************************")
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to SQLite3 database",
		TestFunc: func() {
			dbc2, err2 := sqliteDb.OpenDb()
			defer sqliteDb.CloseDb()
			fmt.Println(dbc2)
			mctest.AssertEquals(t, err2, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	//if dbc != nil || err == nil {
	//	myDb.CloseDb()
	//}
	//if dbc2 != nil || err2 == nil {
	//	sqliteDb.CloseDb()
	//}

	mctest.PostTestResult()
}
