package main

import (
    "payment-service/db"
    "payment-service/kafka"
    "log"
)

func main() {
    // Connect to the database
    db.ConnectDatabase()

    kafka.SetupPaymentProducer("localhost:9092")   

    // Your application logic here
    log.Println("Service is running...")
}
