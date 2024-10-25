package api

import (
    "net/http"
    "order-service/models"
    "order-service/services"
    "order-service/kafka"
    "github.com/gin-gonic/gin"
)

func CreateOrderHandler(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.CreateOrder(&order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := kafka.PublishOrderCreated(order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish order event"})
        return
    }

    c.JSON(http.StatusCreated, order)
}
