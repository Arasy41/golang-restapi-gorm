package entity

type TheaterList struct { 
	ID uint `json:"id" gorm:"primaryKey"`
	TheaterID uint `json:"theaterid"`
	Theather Theather `gorm:"foreignKey:TheaterID"`
	FilmID uint `json:"filmid"`
	Film Film `gorm:"foreignKey:FilmID"`
}

type TheaterDetails struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Kota      string    `json:"kota"`
	Cinema    string    `json:"cinema"`
	Contact   string    `json:"contact"`
	Film      []Film	`json:"film"`
}
