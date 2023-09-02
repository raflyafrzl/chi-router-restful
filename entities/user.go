package entities

type User struct {
	Id          string `gorm:"type:varchar(12);primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(45)" json:"name"`
	Email       string `gorm:"type:varchar(45)" json:"email"`
	Password    string `gorm:"type:varchar(40)" json:"password"`
	PhoneNumber string `gorm:"type:varchar(40)" json:"phone_number"`
	Role        string `gorm:"type:varchar(10)" json:"role"`
	IsVerified  bool   `gorm:"type:boolean" json:"is_verified"`
}
