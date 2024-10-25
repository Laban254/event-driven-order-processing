package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// GetPaymentStatus returns the status of the payment
func GetPaymentStatus(c *gin.Context) {
    // Here you could retrieve the payment status from the database
    c.JSON(http.StatusOK, gin.H{"status": "Payment processed successfully"})
}
