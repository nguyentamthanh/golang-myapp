package models

import "time"

type Product struct {
	Id       uint `json:"id" gorm:"primaryKey"`
	CreateAt time.Time
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
