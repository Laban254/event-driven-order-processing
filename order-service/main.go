package main

import (
    "order-service/db"
    "order-service/api"
    "order-service/kafka"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    db.ConnectDatabase()

    kafka.SetupProducer("localhost:9092")

    router := gin.Default()
    router.POST("/orders", api.CreateOrderHandler)

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
