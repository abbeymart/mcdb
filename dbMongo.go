// @Author: abbeymart | Abi Akindele | @Created: 2020-12-05 | @Updated: 2020-12-05
// @Company: mConnect.biz | @License: MIT
// @Description: mcdb - db connection for MongoDB

package mcdb

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDbConnectionType *mongo.Client

type MongoDbSecureType struct {
	SecureAccess bool
	SecureCert   string
	SecureKey    string
}

type MongoDbConfigType struct {
	Host         string
	Username     string
	Password     string
	DbName       string
	Filename     string
	Location     string
	Port         uint32
	DbType       string
	PoolSize     uint
	SecureOption MongoDbSecureType
	Uri          string
}

type MongoDbConnectOptions map[string]interface{}

type MongoDbConfig struct {
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
	Options MongoDbConnectOptions
}

var (
	dbMg       *mongo.Client
	errMg      error
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func (dbConfig MongoDbConfig) OpenMongoDb() (*mongo.Client, error) {
	// mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
	// mongodb://mongodb0.example.com:27017
	connectionString := fmt.Sprintf("%v://%v:%v", dbConfig.DbType, dbConfig.Host, dbConfig.Port)
	//connectionStringSecure := fmt.Sprintf("%v://%v:%v@%v:%v/%v", dbConfig.DbType, dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	//connectionString := dbConfig.Url
	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	}
	dbMg, errMg = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if errMg != nil {
		errMsg := fmt.Sprintf("Database Connection Error: %v", err)
		return nil, errors.New(errMsg)
	}
	// context options: TODO: review / apply as needed
	ctx, cancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	err = dbMg.Connect(ctx)
	if err != nil {
		errMsg := fmt.Sprintf("Database Connection Error: %v", err)
		return nil, errors.New(errMsg)
	}

	// return db-connection handle
	return dbMg, nil
}

func (dbConfig MongoDbConfig) CloseMongoDb() {
	if dbMg != nil {
		err = dbMg.Disconnect(ctx)
		if err != nil {
			// log error to the console
			fmt.Println(err)
		}
	}
}
