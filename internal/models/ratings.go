package models

import (
	"time"
)

type Rating struct {
	ID int `json:"id"`
	SeriesID int `json:"series_id"`

	Content string `json:"content"`
	StarsQuantity int `json:"stars_quantity"`

	CreatedAt time.Time `json:"created_at"`
}