package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// GetPaymentStatus returns the status of the payment
func GetPaymentStatus(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "Payment processed successfully"})
}
