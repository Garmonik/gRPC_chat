package auth_utils

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database/models"
	"gorm.io/gorm"
)

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
