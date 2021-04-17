package db

import (
	"fmt"
	"notif/model"
	"os"

	"github.com/TomBowyerResearchProject/common/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

//ConnectDB function: Make database connection
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err)
	}

	username := os.Getenv("databaseUser")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")

	//Define DB connection string and connect
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, username, databaseName, password)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		logger.Error(err)
	}

	err = db.AutoMigrate(
		&model.Notification{},
	)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Successfully connected to Database! ALL SYSTEMS ARE GO")
	database = db
}

func GetDB() *gorm.DB {
	return database
}