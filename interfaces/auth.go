package interfaces

import (
	"gochiapp/model"
)

type AuthService interface {
	CompareAndSigned(data model.LoginUserModel) string
	Set(model.UserAuthModel)
	Get(key string) string
	Verified(id string)
}
