package entity

type Theather struct { 
	ID uint `json:"id" gorm:"primaryKeyx"`
	Kota string `json:"kota"`
	Cinema string `json:"cinema"`
	Contact string `json:"contact"`
}
