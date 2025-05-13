package user_rep

import (
	"fmt"
	"store/internal/entities"
	"store/pkg"
)

func (ur *UserRepository) GetPhoneNumber(userid uint) string {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return ""
	}
	return user.Phone
}

func (ur *UserRepository) GetUsername(userid uint) string {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return "یافت نشد"
	}
	if user.FirstName != "" {
		return user.FirstName + " " + user.LastName
	}
	return "ناشناس"
}

func (ur *UserRepository) GetInfo(ID uint) (entities.User, error) {
	var user entities.User
	tx := ur.db.First(&user, ID)
	if tx.Error != nil {
		return user, fmt.Errorf("error getting user")
	}
	user.Password = ""
	user.RefreshToken = ""
	return user, nil
}

func (ur *UserRepository) FillInfo(user entities.User) error {
	var temp entities.User
	ur.db.First(&temp, entities.User{Email: user.Email})
	if temp.ID != user.ID {
		return fmt.Errorf("email exists")
	}
	ur.db.First(&temp, entities.User{Phone: user.Phone})
	if temp.ID != user.ID {
		return fmt.Errorf("phone number exists")
	}
	temp.Phone = user.Phone
	temp.Email = user.Email
	temp.FirstName = user.FirstName
	temp.LastName = user.LastName
	temp.Addresses = user.Addresses
	tx := ur.db.Save(temp)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ur *UserRepository) ResetPassword(userid uint, password string) error {
	pass, _ := pkg.HashPassword(password)
	tx := ur.db.Model(entities.User{}).Where("id = ?", userid).Update("password = ?", pass)
	return tx.Error
}
