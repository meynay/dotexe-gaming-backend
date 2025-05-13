package admin_rep

import (
	"fmt"
	"store/internal/entities"
	"store/pkg"
	"strings"

	"gorm.io/gorm"
)

type AdminRep struct {
	rep *gorm.DB
}

func NewAdminRep(ar *gorm.DB) *AdminRep {
	return &AdminRep{rep: ar}
}

func (ar *AdminRep) GetName(id uint) string {
	admin := entities.Admin{}
	res := ar.rep.First(&admin, id)
	if res.Error != nil {
		return ""
	}
	if admin.FirstName == "" && admin.LastName == "" {
		return "ادمین"
	}
	return admin.FirstName + " " + admin.LastName + " (ادمین)"
}

func (ar *AdminRep) AddAdmin(username, password string) error {
	admin := entities.Admin{}
	username = strings.ToLower(username)
	res := ar.rep.First(&admin, entities.Admin{Username: username})
	if res.Error == nil && admin.Username == username {
		return fmt.Errorf("user exists")
	}
	pass, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}
	admin = entities.Admin{
		Password: pass,
		Username: username,
	}
	tx := ar.rep.Create(&admin)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ar *AdminRep) GetInfo(adminID uint) (entities.Admin, error) {
	admin := entities.Admin{}
	res := ar.rep.First(&admin, adminID)
	if res.Error != nil {
		return admin, res.Error
	}
	admin.Password = ""
	return admin, nil
}

func (ar *AdminRep) FillFields(admin entities.Admin) error {
	var temp entities.Admin
	res := ar.rep.First(&temp, entities.Admin{Phone: admin.Phone})
	if res.Error != nil {
		return res.Error
	}
	if temp.ID != admin.ID {
		return fmt.Errorf("phone number exists")
	}
	tx := ar.rep.Save(&admin)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ar *AdminRep) Login(username, password string) (uint, error) {
	username = strings.ToLower(username)
	admin := entities.Admin{}
	res := ar.rep.First(&admin, entities.Admin{Username: username})
	if res.Error != nil {
		return 0, res.Error
	}
	if pkg.CompareHashAndPassword(admin.Password, password) != nil {
		return 0, fmt.Errorf("wrong password")
	}
	return admin.ID, nil
}
