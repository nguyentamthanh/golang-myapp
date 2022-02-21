package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("faile to connection database", err)
		os.Exit(2)
	}
	log.Println("Connection to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Print("Running Migrations")
	db.AutoMigrate()
	Database = DbInstance{
		Db: db,
	}
}
