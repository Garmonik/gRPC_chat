package auth_utils

import (
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database/models"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/crypto_lib"
	"gorm.io/gorm"
	"time"
)

func GetUserByEmail(email, password string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.Select("id", "email", "password_hash").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateNewSession(userID uint, db *gorm.DB, cfg *config.Config) (string, error) {
	session := models.Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(cfg.SessionTTL),
		IsClosed:  false,
	}

	result := db.Create(&session)
	if result.Error != nil {
		return "", result.Error
	}
	return session.ID.String(), nil
}

func CreateNewUser(email, password, name string, db *gorm.DB, cfg *config.Config) (*models.User, error) {
	hashPassword, err := crypto_lib.HashString(password, cfg)
	if err != nil {
		return nil, err
	}
	user := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashPassword,
		Bio:          "",
	}
}
