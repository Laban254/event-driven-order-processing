package models

import (
    "encoding/json"
)

// Order represents the order model
type Order struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Amount float64 `json:"amount"`
    // Add other relevant fields
}

// ToJSON serializes the order to JSON
func (o Order) ToJSON() string {
    jsonData, err := json.Marshal(o)
    if err != nil {
        return "{}" // Return an empty JSON object on error
    }
    return string(jsonData)
}
