package notification

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotifRepository(db *gorm.DB) interfaces.NotificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (a *notificationRepository) Store(ctx context.Context, data entities.Notification) error {

	err := a.db.WithContext(ctx).Create(data).Error

	if err != nil {
		return err
	}

	return nil

}
