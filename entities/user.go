package entities

type User struct {
	Id          string `gorm:"varchar(12);primaryKey" json:"id"`
	Name        string `gorm:"varchar(45)" json:"name"`
	Email       string `gorm:"varchar(45)" json:"email"`
	Password    string `gorm:"varchar(40)" json:"password"`
	PhoneNumber string `gorm:"varchar(40)" json:"phone_number"`
	Role        string `gorm:"varchar(10)" json:"role"`
}
