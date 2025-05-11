package accessing_database

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GeSessionBySessionID(SessionID uuid.UUID, db *gorm.DB) (*models.Session, error) {
	var session models.Session
	if err := db.Preload("User").Where("id = ?",
		SessionID).
		First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}
