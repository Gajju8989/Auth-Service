package user

import "time"

type User struct {
	ID           string    `gorm:"type:varchar(36);primaryKey"`
	Email        string    `gorm:"type:varchar(255);unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3)"`
}
