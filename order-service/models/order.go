package models

import (
    "encoding/json"
)

// Order represents the order model
type Order struct {
    ID           uint    `json:"id" gorm:"primaryKey"`
    CustomerName string  `json:"customerName"` // Added field for customer name
    Product      string  `json:"product"`      // Added field for product
    Quantity     int     `json:"quantity"`     // Added field for quantity
    Price        float64 `json:"price"`        // Added field for price
    Amount       float64 `json:"amount"`       // Added field for total amount
}

// ToJSON serializes the order to JSON
func (o Order) ToJSON() string {
    jsonData, err := json.Marshal(o)
    if err != nil {
        return "{}" // Return an empty JSON object on error
    }
    return string(jsonData)
}
