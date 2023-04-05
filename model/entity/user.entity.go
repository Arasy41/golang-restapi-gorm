package entity

type User struct {
	ID 			uint 	`json:"id" gorm:"primaryKey"`
	Name 		string 	`json:"name"`
	Email 		string 	`json:"email"`
	Password 	string 	`json:"-" gorm:"index,column:password"`
}