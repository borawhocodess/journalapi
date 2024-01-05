package database

import (
    "journalapi/models"

	"fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// Database represents the database configuration.
type Database struct {
    Host     string
    User     string
    Password string
    DBName   string
    Port     string
}

// ConnectDB initializes a connection to the database and auto-migrates models.
func ConnectDB() *gorm.DB {
    dbConfig := Database{
        Host:     "db", // or your database host
        User:     "root",
        Password: "root",
        DBName:   "journaldb",
        Port:     "5432",
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // AutoMigrate your models here
    db.AutoMigrate(&models.User{}, &models.Entry{}, &models.Journal{})

    return db
}
