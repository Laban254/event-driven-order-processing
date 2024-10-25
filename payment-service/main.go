package main

import (
    "log"
    "net/http"
    "payment-service/db"
    "payment-service/kafka"
    "github.com/gin-gonic/gin"
)

func main() {
    db.ConnectDatabase()

    kafka.SetupPaymentConsumer("localhost:9092")

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "up"})
    })

    if err := r.Run(":8081"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
