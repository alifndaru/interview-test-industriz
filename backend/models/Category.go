package models

import "time"


type Category struct {
	Id 		int64 	`json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string 	`json:"name" gorm:"type:varchar(100);not null;unique"`
	Color  string  `json:"color" gorm:"type:varchar(50)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}