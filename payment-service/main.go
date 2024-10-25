package main

import (
    "log"
    "net/http"
    "payment-service/db"
    "payment-service/kafka"
    "github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    db.ConnectDatabase()

    kafka.SetupPaymentConsumer("localhost:9092")

    // Initialize Gin router
    r := gin.Default()

    // Define routes (if needed, otherwise can be left empty)
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "up"})
    })

    // Start the server on port 8081
    if err := r.Run(":8081"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
