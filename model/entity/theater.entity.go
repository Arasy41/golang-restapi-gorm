package entity

import "time"

type Theather struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Kota      string    `json:"kota"`
	Cinema    string    `json:"cinema"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TheaterDetails struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Kota      string    `json:"kota"`
	Cinema    string    `json:"cinema"`
	Contact   string    `json:"contact"`
	Film      []Film	`json:"film"`
}
