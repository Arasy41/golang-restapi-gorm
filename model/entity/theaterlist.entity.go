package entity

type TheaterList struct { 
	ID uint `json:"id" gorm:"primaryKey"`
	TheaterID uint `json:"theaterid"`
	Theather Theather `gorm:"foreignKey:TheaterID"`
	FilmID uint `json:"filmid"`
	Film Film `gorm:"foreignKey:FilmID"`
}
