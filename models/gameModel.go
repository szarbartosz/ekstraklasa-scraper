package models

import "time"

type Game struct {
	Round    int       `json:"round"`
	DateTime time.Time `json:"dateTime"`
	Host     string    `json:"host"`
	Guest    string    `json:"guest"`
	Result   string    `json:"result"`
}
