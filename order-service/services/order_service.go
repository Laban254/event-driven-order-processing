package services

import (
    "order-service/models"
    "order-service/db"
)

// CreateOrder saves a new order in the database
func CreateOrder(order *models.Order) error {
    // Calculate the amount based on quantity and price
    order.Amount = float64(order.Quantity) * order.Price // Calculate total amount

    result := db.DB.Create(order) // Use the global DB variable
    return result.Error // Return any error that occurs during the create operation
}