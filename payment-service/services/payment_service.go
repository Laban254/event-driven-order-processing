package services

import (
    "payment-service/models"
	"payment-service/db"
    "log"
)

// CreatePayment saves a new payment in the database
func CreatePayment(payment *models.Payment) error {
    result := db.DB.Create(payment) // Use the global DB variable
    return result.Error // Return any error that occurs during the create operation
}

// ProcessPayment handles the business logic for creating a payment
func ProcessPayment(payment *models.Payment) error {
    log.Printf("Processing payment for order ID: %d, Amount: %.2f\n", payment.OrderID, payment.Amount)

    // Save the payment record in the database
    if err := CreatePayment(payment); err != nil {
        return err // Handle error accordingly
    }

    return nil // Indicate success
}
