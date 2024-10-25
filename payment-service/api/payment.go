package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetPaymentStatus(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "Payment processed successfully"})
}
