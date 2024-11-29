package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserName  string    `json:"username" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Status    uint      `json:"status"`
}

//ユーザーが作成されたときに返すレスポンスのデータの型
type UserResponce struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username"`
	Email    string `json:"email" gorm:"unique;not null"`
}
