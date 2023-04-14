package request

import (
	
)

type FilmCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film" validate:"required"`
	Produser	string `json:"produser" validate:"required"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
}

type FilmUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	JenisFilm 	string `json:"jenis_film"`
	Produser	string `json:"produser"`
	Sutradara	string `json:"sutradara"`
	Penulis		string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts		string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
	Like 		uint   `json:"like"`
}

type FilmLikeUpdateRequest struct {
	Like 		uint 	`json:"like"`
}

type UserCreateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Name 		string `json:"name" validate:"required"` 
	Email 		string 	`json:"email" validate:"required"`
	Password 	string 	`json:"password" validate:"required"`
}

type UserEmailRequest struct {
	Email 		string 	`json:"email" validate:"required"`
}

type TheaterCreateRequest struct {
	Kota string `json:"kota"`
	Cinema string `json:"cinema"`
	Contact string `json:"contact"`
}

type TheaterListCreateRequest struct {
	TheaterID uint `json:"theaterid"`
	FilmID uint `json:"filmid"`
}

type TheaterUpdateRequest struct {
	Kota string `json:"kota"`
	Cinema string `json:"cinema"`
	Contact string `json:"contact"`
}

type CommentCreateRequest struct {
	FilmID    uint      `json:"film_id" validate:"required"`
	Comment   string    `json:"comment" validate:"required"`
}

