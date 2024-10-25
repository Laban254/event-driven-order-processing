package models

import (
    "encoding/json"
)

type Order struct {
    ID           uint    `json:"id" gorm:"primaryKey"`
    CustomerName string  `json:"customerName"` 
    Product      string  `json:"product"`     
    Quantity     int     `json:"quantity"`    
    Price        float64 `json:"price"`       
    Amount       float64 `json:"amount"`     
}

func (o Order) ToJSON() string {
    jsonData, err := json.Marshal(o)
    if err != nil {
        return "{}" 
    }
    return string(jsonData)
}
