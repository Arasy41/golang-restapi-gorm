package entity

type CommentList struct {
	Id		uint	`json:"id"`
	FilmId	uint 	`json:"filmid"`
	Film    Film	`gorm:"foreignKey:FilmId"`
	CommentId uint  `json:"commentid"`
	Comment	Comment	`gorm:"foreignKey:CommentId"`
}