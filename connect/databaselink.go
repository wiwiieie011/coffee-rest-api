package connect

import (
	"dzabrail/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectBase() {
	db, err := gorm.Open(sqlite.Open("coffe.db"), &gorm.Config{})

	if err != nil {
		panic("fail connection db")
	}

	err = db.AutoMigrate(&models.Drink{})

	if err !=nil{
		panic("failed to migrate database")
	}

	DB = db



}
