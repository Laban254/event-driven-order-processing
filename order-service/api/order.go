package api

import (
    "net/http"
    "order-service/models"
    "order-service/services"
    "order-service/kafka"
    "github.com/gin-gonic/gin"
)

// CreateOrderHandler handles the HTTP request for creating an order
func CreateOrderHandler(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create the order in the database
    if err := services.CreateOrder(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Publish the order created event to Kafka
    if err := kafka.PublishOrderCreated(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish order event"})
        return
    }

    c.JSON(http.StatusCreated, order)
}
