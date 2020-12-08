// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: go: mConnect

package mcdb

import (
	"fmt"
	"testing"
)
import "github.com/abbeymart/mctestgo"

func TestDbMongo(t *testing.T) {
	// test-data: db-configuration settings
	myDb := MongoDbConfig{
		DbType:   "postgres",
		Host:     "localhost",
		Username: "abbeymart",
		Password: "ab12testing",
		Port:     27017,
		DbName:   "mcdev",
		Filename: "testdb.db",
		PoolSize: 20,
		Url:      "mongodb://localhost:27017/mcdev",
	}
	myDb.Options = MongoDbConnectOptions{}

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the mongoDB:",
		TestFunc: func() {
			dbc, err := myDb.OpenMongoDb()
			defer myDb.CloseMongoDb()
			fmt.Println(dbc)
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			//mctest.AssertEquals(t, req.Message, res.Message, "response-message should be: "+res.Message)
		},
	})

	mctest.PostTestResult()
}
