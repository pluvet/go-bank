package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pluvet/go-bank/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetMockDB() sqlmock.Sqlmock {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})

	config.DB = db

	return mock
}
