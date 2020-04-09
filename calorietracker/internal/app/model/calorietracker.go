package model

import "time"

type Config struct {
	Birthday time.Time `json:"birthday"`
	Height   float     `json:"height"`
}
type Weight struct {
	Id     string    `json:"id"`
	Date   time.Time `json:"date"`
	Weight float64   `json:"weight"`
}
type Weights struct {
	Weight []Weight `json:"weights"`
}
type Entry struct {
	Date     time.Time `json:"date"`
	Food     string    `json:"food"`
	Calories int       `json:"calories"`
}
