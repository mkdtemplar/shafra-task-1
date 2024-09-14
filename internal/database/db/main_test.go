package db

import (
	"fmt"
	"log"
	"os"
	"shafra-task-1/utils"
	"testing"

	"gorm.io/gorm"
)

var conn *gorm.DB

func TestMain(m *testing.M) {

	config, err := utils.LoadConfig("../../..")
	conn = ConnectToPostgres(config)

	if err != nil {
		fmt.Println("Cannot connect to  database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}
	os.Exit(m.Run())
}
