package repositories

import (
	"errors"

	"github.com/pluvet/go-bank/app/config"
	"github.com/pluvet/go-bank/app/models"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	CreateAccount(int) (*int, error)
	FindAccount(int) (*models.Account, error)
	UpdateAccount(account *models.Account) error
}

type AccountRepository struct{}

func (a *AccountRepository) CreateAccount(userID int) (*int, error) {
	var account models.Account
	account.Balance = 0
	account.UserID = uint64(userID)
	result := config.DB.Create(&account)

	if result.Error != nil {
		err := new(ErrorCreatingRecordInDB)
		err.Model = "account"
		return nil, err
	}

	return &account.ID, nil
}

func (a *AccountRepository) FindAccount(accountID int) (*models.Account, error) {
	var account *models.Account
	dbErr := config.DB.Where("id = ?", accountID).First(&account).Error
	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		err := new(ErrorFindingOneRecordInDB)
		err.Model = "account"
		err.ID = accountID
		return nil, err
	}
	return account, nil
}

func (a *AccountRepository) UpdateAccount(account *models.Account) error {
	result := config.DB.Save(account)
	if result.Error != nil {
		err := new(ErrorUpdatingRecordInDB)
		err.Model = "account"
		err.ID = account.ID
		return err
	}
	return nil
}
