package user_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uu *UserUsecase) GetInfo(ID primitive.ObjectID) (entities.User, error) {
	return uu.userRep.GetInfo(ID)
}

func (uu *UserUsecase) FillInfo(user entities.User) error {
	return uu.userRep.FillInfo(user)
}

func (uu *UserUsecase) ResetPassword(userid primitive.ObjectID, password string) error {
	return uu.userRep.ResetPassword(userid, password)
}
