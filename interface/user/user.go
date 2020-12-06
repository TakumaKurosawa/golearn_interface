package user

import "golearninterface/model/user/entity"

type UserInterface interface {
	Print()
	UpdateBaseData(entity.UserData)
}
