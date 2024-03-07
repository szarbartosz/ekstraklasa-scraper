package models

type Game struct {
	Round    int    `json:"round"`
	DateTime string `json:"dateTime"`
	Host     string `json:"host"`
	Guest    string `json:"guest"`
	Result   string `json:"result"`
}
