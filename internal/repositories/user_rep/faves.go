package user_rep

import (
	"fmt"
	"store/internal/entities"
	"store/pkg"
)

func (ur *UserRepository) AddToFaves(productid, userid uint) error {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	user.Faves = append(user.Faves, productid)
	return nil
}

func (ur *UserRepository) DeleteFromFaves(productid, userid uint) error {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	for index, pid := range user.Faves {
		if pid == productid {
			user.Faves = append(user.Faves[:index], user.Faves[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("product was not in user faves")
}

func (ur *UserRepository) CheckFave(productid, userid uint) error {
	var user entities.User
	res := ur.db.First(&user, userid)
	if res.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	if pkg.Exists(productid, user.Faves) {
		return nil
	}
	return fmt.Errorf("product doesn't exist in faves")
}

func (ur *UserRepository) GetFaves(userid uint) []uint {
	var u entities.User
	res := ur.db.First(&u, userid)
	if res.Error != nil {
		return []uint{}
	}
	return u.Faves
}
