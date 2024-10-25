package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "fmt" 
    "order-service/config" 
)

var DB *gorm.DB

func ConnectDatabase() {
    cfg := config.LoadConfig()

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connection established")

    MigrateDatabase() 
}

