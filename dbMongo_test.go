// @Author: abbeymart | Abi Akindele | @Created: 2020-12-04 | @Updated: 2020-12-04
// @Company: mConnect.biz | @License: MIT
// @Description: db-mongodb testing

package mcdb

import (
	"fmt"
	"testing"
)
import "github.com/abbeymart/mctestgo"

func TestDbMongo(t *testing.T) {
	// test-data: db-configuration settings
	myDb := MongoDbConfig{
		DbType:   "mongodb",
		Host:     "localhost",
		Username: "abbeymart",
		Password: "ab12testing",
		Port:     27017,
		DbName:   "mcdev",
		Filename: "testdb.db",
		PoolSize: 20,
		Url:      "mongodb://localhost:27017",
	}
	myDb.Options = MongoDbConnectOptions{}

	mctest.McTest(mctest.OptionValue{
		Name: "should successfully connect to the mongoDB Host/Server:",
		TestFunc: func() {
			mgServer, _ := myDb.OpenMongoDb()
			mgServerDb := mgServer.Database(myDb.DbName)
			defer myDb.CloseMongoDb()
			fmt.Println(mgServer)
			fmt.Println("********************************")
			fmt.Println(mgServerDb.Name())
			fmt.Println("*********")
			mctest.AssertEquals(t, err, nil, "response-code should be: nil")
			mctest.AssertEquals(t, mgServerDb.Name(), myDb.DbName, "response-message should be: "+myDb.DbName)
		},
	})

	mctest.PostTestResult()
}
