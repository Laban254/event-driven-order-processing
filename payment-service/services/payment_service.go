package services

import (
    "fmt"
    "payment-service/models"
    "payment-service/db"
    "log"
)

func CreatePayment(payment *models.Payment) error {
    result := db.DB.Create(payment) 
    if err := result.Error; err != nil {
        log.Printf("Error creating payment: %v", err) 
        return err 
    }
    return nil 
}

func ProcessPayment(payment *models.Payment) error {
    log.Printf("Processing payment for Order ID: %d, Amount: %.2f\n", payment.OrderID, payment.Amount)

    // Validate the payment details before proceeding
    if payment.OrderID == 0 {
        log.Println("Invalid OrderID: cannot process payment")
        return fmt.Errorf("invalid OrderID: %d", payment.OrderID) 
    }

    if err := CreatePayment(payment); err != nil {
        return err 
    }

    payment.Status = "completed"

    log.Printf("Successfully processed payment for Order ID: %d", payment.OrderID)

    return nil 
}
