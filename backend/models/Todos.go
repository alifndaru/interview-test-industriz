package models

import "time"


type Todos struct {
	Id 			int64 		`json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     	`json:"title" gorm:"type:varchar(255);not null"`	
	Description	string     	`json:"description"`
	Completed	  	bool     `json:"completed" gorm:"type:boolean;not null;default:false"`
	CategoryID  int64      `json:"category_id"`
	Category    *Category   `json:"category" gorm:"foreignKey:CategoryID;references:Id"`
	Priority    string     `gorm:"type:varchar(10)" json:"priority"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}