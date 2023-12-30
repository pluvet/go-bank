package config 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )

fun Connect() {
	url := "postgres://postgres:123456@pgdb:5432/bank"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
}
  
  
  