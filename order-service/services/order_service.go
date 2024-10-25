package services

import (
    "order-service/models"
    "order-service/db"
)

func CreateOrder(order *models.Order) error {
    order.Amount = float64(order.Quantity) * order.Price // Calculate total amount

    result := db.DB.Create(order)
    return result.Error 
}