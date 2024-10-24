package api

import (
    "net/http"
    "order-service/models"
    "order-service/services"
    "github.com/gin-gonic/gin"
)

// CreateOrderHandler handles the HTTP request for creating an order
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

    c.JSON(http.StatusCreated, order)
}
