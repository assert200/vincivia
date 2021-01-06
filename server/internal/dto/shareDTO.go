package dto

import "time"

// Share form DTO
type share struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Volume      int       `json:"volume"`
	Timestamp   time.Time `json:"timestamp"`
}
