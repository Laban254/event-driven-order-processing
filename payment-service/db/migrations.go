package db

import "payment-service/models"

func MigrateDatabase() error {
    return DB.AutoMigrate(&models.Payment{})
}
