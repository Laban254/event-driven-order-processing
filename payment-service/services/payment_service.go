package services

import (
    "fmt"
    "payment-service/models"
    "payment-service/db"
    "log"
)

// CreatePayment saves a new payment in the database
func CreatePayment(payment *models.Payment) error {
    result := db.DB.Create(payment) // Use the global DB variable
    if err := result.Error; err != nil {
        log.Printf("Error creating payment: %v", err) // Log the error for debugging
        return err // Return any error that occurs during the create operation
    }
    return nil // Indicate success
}

// ProcessPayment handles the business logic for creating a payment
func ProcessPayment(payment *models.Payment) error {
    // Log the initial processing of the payment
    log.Printf("Processing payment for Order ID: %d, Amount: %.2f\n", payment.OrderID, payment.Amount)

    // Validate the payment details before proceeding
    if payment.OrderID == 0 {
        log.Println("Invalid OrderID: cannot process payment")
        return fmt.Errorf("invalid OrderID: %d", payment.OrderID) // Return an error if OrderID is invalid
    }

    // Save the payment record in the database
    if err := CreatePayment(payment); err != nil {
        return err // Handle error accordingly
    }

    // Mark payment as completed
    payment.Status = "completed"

    // Log the successful payment processing
    log.Printf("Successfully processed payment for Order ID: %d", payment.OrderID)

    return nil // Indicate success
}
