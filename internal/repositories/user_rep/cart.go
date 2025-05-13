package user_rep

import (
	"fmt"
	"store/internal/entities"
)

func (ur *UserRepository) AddToCart(productid, userid uint) error {
	item := entities.Item{
		ProductID: productid,
		Count:     1,
	}
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	user.Cart = append(user.Cart, item)
	return nil
}

func (ur *UserRepository) DeleteFromCart(productid, userid uint) error {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	for index, c := range user.Cart {
		if c.ProductID == productid {
			user.Cart = append(user.Cart[:index], user.Cart[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("couldn't find item in user cart")
}

func (ur *UserRepository) IsInCart(productid, userid uint) (int, error) {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return 0, fmt.Errorf("couldn't find user")
	}
	for _, item := range user.Cart {
		if item.ProductID == productid {
			return item.Count, nil
		}
	}
	return 0, fmt.Errorf("couldn't find product in user cart")
}

func (ur *UserRepository) IncreaseInCart(productid, userid uint) error {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	for index, item := range user.Cart {
		if item.ProductID == productid {
			item.Count++
			user.Cart = append(user.Cart[:index], user.Cart[index+1:]...)
			user.Cart = append(user.Cart, item)
			return nil
		}
	}
	return fmt.Errorf("couldn't find product in user cart")
}

func (ur *UserRepository) DecreaseInCart(productid, userid uint) error {
	var user entities.User
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return fmt.Errorf("couldn't find user")
	}
	for index, item := range user.Cart {
		if item.ProductID == productid {
			item.Count--
			user.Cart = append(user.Cart[:index], user.Cart[index+1:]...)
			if item.Count != 0 {
				user.Cart = append(user.Cart, item)
			}
			return nil
		}
	}
	return fmt.Errorf("couldn't find product in user cart")
}

func (ur *UserRepository) GetCart(userid uint) ([]entities.Item, error) {
	user := entities.User{}
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return user.Cart, fmt.Errorf("couldn't find user")
	}
	return user.Cart, nil
}

func (ur *UserRepository) FinalizeCart(userid uint) ([]entities.Item, error) {
	user := entities.User{}
	tx := ur.db.First(&user, userid)
	if tx.Error != nil {
		return user.Cart, fmt.Errorf("couldn't find user")
	}
	cart := user.Cart
	user.Cart = []entities.Item{}
	ur.db.Save(&user)
	return cart, nil
}
