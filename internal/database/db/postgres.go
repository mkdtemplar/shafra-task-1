package db

import (
	"fmt"
	"log"
	"shafra-task-1/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn = &PostgresDB{}

func ConnectToPostgres(config utils.Config) *gorm.DB {

	dbConn, err := gorm.Open(postgres.Open(config.DbSource), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to  database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	return dbConn
}

func GetDb() *gorm.DB {
	return Conn.DB
}
