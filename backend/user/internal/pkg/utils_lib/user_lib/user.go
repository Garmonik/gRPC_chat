package user_lib

import (
	"context"
	"fmt"
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

func GetUserByEmail(ctx context.Context, email string, db *gorm.DB) (*models.User, error) {
	var user models.User
	if err := db.WithContext(ctx).Select("id", "email", "password_hash").
		Where("email = ?", email).
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

func UpdateUser(ctx context.Context, userID int64, email, bio string, db *gorm.DB) (int, error) {
	result := db.WithContext(ctx).Model(&models.User{}).
		Where("id = ?", userID).
		Update("bio", bio).
		Update("email = ?", email)
	if result.Error != nil {
		return 1, fmt.Errorf("database error: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return 2, fmt.Errorf("user not found")
	}

	return 0, nil
}

func GetUserList(ctx context.Context, UserId int64, orderBy string, asc bool, search string, db *gorm.DB) ([]*models.User, error) {
	query := db.WithContext(ctx).Model(&models.User{}).
		Where("id != ?", UserId)

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	sortField := "created_at"
	switch orderBy {
	case "name":
		sortField = "name"
	case "created":
		sortField = "created_at"
	}

	sortOrder := "ASC"
	if !asc {
		sortOrder = "DESC"
	}

	query = query.Order(sortField + " " + sortOrder)

	var users []*models.User
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
