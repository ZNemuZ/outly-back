package model

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserId    int       `json:"user_id" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null;type:varchar(50)"`
	Content   string    `json:"content" gorm:"type:varchar(300)"`
	NiceCount int       `json:"nice_count" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
}

type PostResponce struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null;type:varchar(50)"`
	Content   string    `json:"content" gorm:"type:varchar(300)"`
	NiceCount int       `json:"nice_count" gorm:"not null;default:0"`
	CreatedAt time.Time `json:"created_at"`
}
