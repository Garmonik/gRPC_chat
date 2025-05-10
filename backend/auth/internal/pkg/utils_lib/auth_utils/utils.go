package auth_utils

import (
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database/models"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/crypto_lib"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GetUserByEmail(email string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.Select("id", "email", "password_hash").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(name string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.Select("id", "name", "email", "password_hash").
		Where("name = ?", name).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id uint, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.Select("id", "name", "email", "password_hash").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckSessionID(sessionUuid string, userID uint, db *gorm.DB) (*models.Session, error) {
	sessionID, err := uuid.Parse(sessionUuid)
	if err != nil {
		return nil, fmt.Errorf("invalid session UUID: %v", err)
	}
	var session models.Session
	if err := db.Select("id", "user_id", "is_closed").
		Where("id = ? AND user_id  = ? AND is_closed = false", sessionID, userID).
		First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func CloseSession(sessionUuid string, userID uint, db *gorm.DB) error {
	sessionID, err := uuid.Parse(sessionUuid)
	if err != nil {
		return fmt.Errorf("invalid session UUID: %v", err)
	}

	result := db.Model(&models.Session{}).
		Where("id = ? AND user_id = ?", sessionID, userID).
		Update("is_closed", true)
	if result.Error != nil {
		return fmt.Errorf("database error: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("session not found or already closed")
	}

	return nil
}

func CreateNewSession(userID uint, ipAddress string, db *gorm.DB, cfg *config.Config) (string, error) {
	session := models.Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(cfg.SessionTTL),
		IsClosed:  false,
		IPAddress: ipAddress,
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
	if err := db.Create(&user).Error; err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &user, nil
}
