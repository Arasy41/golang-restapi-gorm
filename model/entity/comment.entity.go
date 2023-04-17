package entity

import "time"

type Comment struct {
	Id        uint      `json:"id"`
	FilmID    uint      `json:"film_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
