package storage

import (
	"log"
	"time"

	"github.com/GaijinZ/user-api/src/rest_api/config"
	"github.com/GaijinZ/user-api/src/rest_api/model"
	"github.com/gocql/gocql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Session *gocql.Session

func ConnectGorm(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	model.Migrate(DB)

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func ConnectCassandra() *gocql.Session {
	var err error

	cluster := gocql.NewCluster("192.168.33.2")
	cluster.Keyspace = "usersapi"
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = time.Second * 10
	Session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	return Session
}
