package initializers

import (
	"fmt"
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=localhost user=postgres password=kellyakacj dbname=trekkers port=5432")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("Connection to DB was Successful🎉"+ db.Name())
	if err != nil {
		panic("Failed to connect to the database!")
	}

	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Step{})
	DB = db

}
