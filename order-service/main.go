package main

import (
    "order-service/db"
    "order-service/api"
    "order-service/kafka"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    // Connect to the database
    db.ConnectDatabase()

    // Set up Kafka producer
    kafka.SetupProducer("localhost:9092")

    // Set up Gin router
    router := gin.Default()
    router.POST("/orders", api.CreateOrderHandler)

    // Start the Gin server on port 8080
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
