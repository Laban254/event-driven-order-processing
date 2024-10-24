package db

import "payment-service/models"

// Migrate performs database migrations for the models
func MigrateDatabase() error {
    return DB.AutoMigrate(&models.Payment{})
}
