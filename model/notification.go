package model

type CreateNotificationModel struct {
	UserId  string `json:"user_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
