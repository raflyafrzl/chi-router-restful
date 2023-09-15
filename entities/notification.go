package entities

type Notification struct {
	Id         string `gorm:"type:varchar(10);primaryKey" json:"id"`
	UserId     string `gorm:"type:varchar(14)" json:"user_id"`
	Message    string `gorm:"type:text" json:"message"`
	MarkAsRead bool   `gorm:"type:boolean" json:"mark_as_read"`
}
