package models

import (
    "encoding/json"
    // "payment-service/db"
)

// Payment represents the payment model
type Payment struct {
    ID      uint    `json:"id" gorm:"primaryKey"`
    OrderID uint    `json:"order_id"`
    Amount  float64 `json:"amount"`
    Status  string  `json:"status"`
}

// ToJSON serializes the payment to JSON
func (p Payment) ToJSON() string {
    jsonData, err := json.Marshal(p)
    if err != nil {
        return "{}" // Return an empty JSON object on error
    }
    return string(jsonData)
}
