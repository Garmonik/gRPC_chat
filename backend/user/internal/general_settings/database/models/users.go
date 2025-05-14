package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:255"`
	Email        string    `gorm:"not null;uniqueIndex;size:255"`
	PasswordHash string    `gorm:"not null;size:255"`
	Bio          string    `gorm:"type:text"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (User) TableName() string {
	return "userprofile"
}

type Session struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	User      User      `gorm:"foreignKey:UserID"`
	UserID    uint
	CreatedAt time.Time `gorm:"autoCreateTime"`
	ExpiresAt time.Time `gorm:"index"`
	IsClosed  bool      `gorm:"default:false"`
	IPAddress string    `gorm:"size:255"`
}

func (Session) TableName() string {
	return "session"
}

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) Close() {
	s.IsClosed = true
}
