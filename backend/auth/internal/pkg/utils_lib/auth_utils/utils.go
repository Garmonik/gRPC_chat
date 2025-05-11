package auth_utils

import (
	"context"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database/models"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/crypto_lib"
	authv1 "github.com/Garmonik/gRPC_chat/backend/protos/gen/go/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GetUserByEmail(ctx context.Context, email string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Select("id", "email", "password_hash").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(ctx context.Context, name string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Select("id", "name", "email", "password_hash").
		Where("name = ?", name).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(ctx context.Context, id uint, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Select("id", "name", "email", "password_hash").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckSessionID(ctx context.Context, sessionUuid string, userID uint, db *gorm.DB) (*models.Session, error) {
	sessionID, err := uuid.Parse(sessionUuid)
	if err != nil {
		return nil, fmt.Errorf("invalid session UUID: %v", err)
	}
	var session models.Session
	if err := db.WithContext(ctx).Select("id", "user_id", "is_closed").
		Where("id = ? AND user_id  = ? AND is_closed = false", sessionID, userID).
		First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func CloseSession(ctx context.Context, sessionUuid string, userID uint, db *gorm.DB) error {
	sessionID, err := uuid.Parse(sessionUuid)
	if err != nil {
		return fmt.Errorf("invalid session UUID: %v", err)
	}

	result := db.WithContext(ctx).Model(&models.Session{}).
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

func CreateNewSession(ctx context.Context, userID uint, ipAddress string, db *gorm.DB, cfg *config.Config) (string, error) {
	session := models.Session{
		UserID:    userID,
		ExpiresAt: time.Now().Add(cfg.SessionTTL),
		IsClosed:  false,
		IPAddress: ipAddress,
	}

	result := db.WithContext(ctx).Create(&session)
	if result.Error != nil {
		return "", result.Error
	}
	return session.ID.String(), nil
}

func CreateNewUser(ctx context.Context, email, password, name string, db *gorm.DB, cfg *config.Config) (*models.User, error) {
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
	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &user, nil
}

func GetActiveSession(ctx context.Context, userID uint64, db *gorm.DB) ([]models.Session, error) {
	var sessions []models.Session
	err := db.WithContext(ctx).Model(&models.Session{}).
		Preload("User").
		Where("user_id = ? AND is_closed = ? AND expires_at > ?",
			userID,
			false,
			time.Now(),
		).Find(&sessions).Error

	if err != nil {
		return nil, fmt.Errorf("GetActiveSession: %w", err)
	}
	return sessions, nil
}

func ConvertSessionToPB(session models.Session) *authv1.Session {
	return &authv1.Session{
		Id: session.ID.String(),
		User: &authv1.User{
			Id:    uint64(session.User.ID),
			Name:  session.User.Name,
			Email: session.User.Email,
		},
		IpAddress: session.IPAddress,
		CreatedAt: session.CreatedAt.Format(time.RFC3339),
		ExpiresAt: session.ExpiresAt.Format(time.RFC3339),
		IsClosed:  session.IsClosed,
	}
}
