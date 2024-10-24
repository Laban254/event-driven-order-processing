package services

import (
    "order-service/models"
    "order-service/db"
)

// CreateOrder saves a new order in the database
func CreateOrder(order *models.Order) error {
    result := db.DB.Create(order) // Use the global DB variable
    return result.Error // Return any error that occurs during the create operation
}
