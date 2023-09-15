package notification

import (
	"context"
	"gochiapp/entities"
	"gochiapp/interfaces"
	"gochiapp/model"
	"gochiapp/utils"
	"time"
)

type notificationService struct {
	repo interfaces.NotificationRepository
}

func NewNotifService(r *interfaces.NotificationRepository) interfaces.NotificationService {
	return &notificationService{
		repo: *r,
	}
}

func (n *notificationService) Create(payload model.CreateNotificationModel) entities.Notification {

	utils.Validate[model.CreateNotificationModel](payload)

	var data entities.Notification = entities.Notification{
		Id:      utils.GetRandomID(10),
		UserId:  payload.UserId,
		Message: payload.Message,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	var err error

	defer cancel()

	err = n.repo.Store(ctx, data)

	if err != nil {
		utils.ErrorResponseWeb(err, 400)
	}

	return data

}
