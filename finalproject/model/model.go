package model

import "time"

type Comment struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Date time.Time `json:"date"`
}
