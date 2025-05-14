package user_lib

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database/models"
	"gorm.io/gorm"
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

func GetUserByID(ctx context.Context, id uint64, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Select("id", "name", "email", "password_hash").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
