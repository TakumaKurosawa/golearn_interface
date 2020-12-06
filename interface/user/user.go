package user

import "golearninterface/model/user/entity"

type UserInterface interface {
	PrintMyData()
	UpdateBaseData(entity.UserData)
}
