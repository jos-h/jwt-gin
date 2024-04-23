package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	DbName    string
	DbHost    string
	DbUser    string
	DbPort    string
	DbSSLMode string
	DbPass    string
}

func createDatabase(config Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass)

	// Connect to PostgreSQL server without specifying a database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL server for creating database: %v", err)
	}

	// Check if database is already created, if yes then skip else create
	db_check := db.Exec(fmt.Sprintf("SELECT datname FROM pg_catalog.pg_database WHERE datname='%s'", config.DbName))
	if db_check.Error != nil{
		// Create the database
		result := db.Exec(fmt.Sprintf("CREATE DATABASE %s", config.DbName))
		if result.Error != nil {
			return fmt.Errorf("failed to create database: %v", result.Error)
		}
	}
	return nil
}
func SetupDB() {
	// Load DB creds from env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error in loading .env file")
	}

	config := Config{
		DbHost:    os.Getenv("DB_HOST"),
		DbUser:    os.Getenv("DB_USER"),
		DbPass:    os.Getenv("DB_PASSWORD"),
		DbPort:    os.Getenv("DB_PORT"),
		DbName:    os.Getenv("DB_NAME"),
		DbSSLMode: os.Getenv("DB_SSLMODE"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbUser, config.DbPass, config.DbPort, config.DbName, config.DbSSLMode)

	err = createDatabase(config)

	if err != nil {
    	log.Fatal("failed to create database", err.Error())
    }

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Connection error", err.Error())
	}
	DB.AutoMigrate(&User{})
}
