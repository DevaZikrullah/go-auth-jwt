package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("parsa:password@/go_auth"), &gorm.Config{})

	if err != nil {
		panic("could connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
