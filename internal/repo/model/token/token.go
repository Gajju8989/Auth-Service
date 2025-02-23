package token

import (
	"github/com/Gajju8989/Auth_Service/internal/repo/model/user"
	"time"
)

type AccessToken struct {
	ID        string     `gorm:"type:varchar(36);primaryKey"`
	UserID    string     `gorm:"type:varchar(36);not null;index"`
	User      user.User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	TokenHash string     `gorm:"type:varchar(255);not null"`
	ExpiresAt time.Time  `gorm:"not null"`
	CreatedAt time.Time  `gorm:"default:current_timestamp"`
	RevokedAt *time.Time `gorm:"default:null"`
}
