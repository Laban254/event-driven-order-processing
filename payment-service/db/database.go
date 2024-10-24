package db

import (
    "log"
    "fmt" // Import the fmt package for string formatting
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "payment-service/config" 
)

var DB *gorm.DB

// ConnectDatabase establishes a connection to the PostgreSQL database
func ConnectDatabase() {
    // Load configuration values from the config package
    cfg := config.LoadConfig()

    // Construct the Data Source Name (DSN) using environment variables
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

    // Establish the database connection
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connection established")

    // Perform database migrations
    MigrateDatabase() // Call to the migration function
}

