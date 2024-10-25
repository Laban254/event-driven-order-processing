package db

import "order-service/models"

func MigrateDatabase() error {
    return DB.AutoMigrate(&models.Order{})
}
