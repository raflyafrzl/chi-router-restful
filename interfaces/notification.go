package interfaces

import (
	"context"
	"gochiapp/entities"
	"gochiapp/model"
)

type NotificationService interface {
	Create(payload model.CreateNotificationModel) entities.Notification
}
type NotificationRepository interface {
	Store(ctx context.Context, data entities.Notification) error
}
