package user_usecase

import (
	"store/internal/entities"
)

func (uu *UserUsecase) GetInfo(ID uint) (entities.User, error) {
	return uu.userRep.GetInfo(ID)
}

func (uu *UserUsecase) FillInfo(user entities.User) error {
	return uu.userRep.FillInfo(user)
}

func (uu *UserUsecase) ResetPassword(userid uint, password string) error {
	return uu.userRep.ResetPassword(userid, password)
}
