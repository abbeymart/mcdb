// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcdb

import (
	"fmt"
	"testing"
)
import "github.com/abbeymart/mctestgo"

func TestSetCache(t *testing.T) {
	// test-data: db-configuration settings
	myDb := DbConfig{
		DbType: "postgres",
	}
	myDb.Host = "localhost"
	myDb.Username = "postgres"
	myDb.Password = "ab12trust"
	myDb.Port = 5432
	myDb.DbName = "mcdev"
	myDb.Filename = "testdb.db"
	myDb.PoolSize = 20
	myDb.Uri = "localhost:5432"
	myDb.Options = DbConnectOptions{}

	sqliteDb := DbConfig{
		DbType: "sqlite3",
	}
	sqliteDb.Filename = "testdb.db"

	var (
		dbc, dbc2 DbConnectionType
		err, err2 error
	)

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the database",
		TestFunc: func() {
			dbc, err := myDb.OpenDb()
			fmt.Println(dbc)
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to SQLite3 database",
		TestFunc: func() {
			dbc2, err2 := sqliteDb.OpenDb()
			fmt.Println(dbc2)
			mctest.AssertEquals(t, err2, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	if dbc != nil || err == nil {
		myDb.CloseDb()
	}
	if dbc2 != nil || err2 == nil {
		sqliteDb.CloseDb()
	}

	mctest.PostTestResult()
}
