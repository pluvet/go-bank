package services

import (
	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
)

func CreateUser(user *models.User) error {

	result := config.DB.Create(&user)
	if result.Error != nil {
		err := new(CreateError)
		err.Model = "user"
		return err
	}

	return nil
}
