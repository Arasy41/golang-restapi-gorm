package entity

import "time"

type Film struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	JenisFilm string    `json:"jenis_film"`
	Produser  string    `json:"produser"`
	Sutradara string    `json:"sutradara"`
	Penulis   string    `json:"penulis"`
	Produksi  string    `json:"produksi"`
	Casts     string    `json:"casts"`
	Sinopsis  string    `json:"sinopsis"`
	Like      uint      `json:"like"`
	Cover     string	`json:"cover"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TheaterId struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TheaterId uint 		`json:"theater_id"`
	Name      string    `json:"name"`
	JenisFilm string    `json:"jenis_film"`
	Produser  string    `json:"produser"`
	Sutradara string    `json:"sutradara"`
	Penulis   string    `json:"penulis"`
	Produksi  string    `json:"produksi"`
	Casts     string    `json:"casts"`
	Sinopsis  string    `json:"sinopsis"`
	Like      uint      `json:"like"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FilmComment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	JenisFilm string    `json:"jenis_film"`
	Produser  string    `json:"produser"`
	Sutradara string    `json:"sutradara"`
	Penulis   string    `json:"penulis"`
	Produksi  string    `json:"produksi"`
	Casts     string    `json:"casts"`
	Sinopsis  string    `json:"sinopsis"`
	Like      uint      `json:"like"`
	Comment   string 	`json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}