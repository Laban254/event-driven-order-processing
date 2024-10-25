package models

import (
    "encoding/json"
)


type Payment struct {
    ID      uint    `json:"id" gorm:"primaryKey"`
    OrderID uint    `json:"order_id"`
    Amount  float64 `json:"amount"`
    Status  string  `json:"status"` 
}

func (p Payment) ToJSON() string {
    jsonData, err := json.Marshal(p)
    if err != nil {
        return "{}" 
    }
    return string(jsonData)
}
