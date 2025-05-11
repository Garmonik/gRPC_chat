package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name         string    `gorm:"size:255" json:"name,omitempty"`
	Email        string    `gorm:"not null;uniqueIndex;size:255" json:"email,omitempty"`
	PasswordHash string    `gorm:"not null;size:255" json:"-"`
	Bio          string    `gorm:"type:text" json:"bio,omitempty"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
}

func (User) TableName() string {
	return "userprofile"
}

type Session struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	UserID    uint      `json:"userId,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	ExpiresAt time.Time `gorm:"index" json:"expiresAt,omitempty"`
	IsClosed  bool      `gorm:"default:false" json:"isClosed,omitempty"`
	IPAddress string    `gorm:"size:255" json:"ipAddress,omitempty"`
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
